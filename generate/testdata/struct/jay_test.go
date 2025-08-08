package main

import (
	"testing"

	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
)

func TestFuzz_1(t *testing.T) {
	var expected, actual One
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, One{}, expected)
	require.Equal(t, One{}, actual)

	expected = One{
		One: rando.Bool(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, One{}, expected)
	// require.NotEqual(t, One{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_2(t *testing.T) {
	var expected, actual Two
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Two{}, expected)
	require.Equal(t, Two{}, actual)

	expected = Two{
		One: One{One: rando.Bool()},
		Two: rando.BoolsN(7),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Two{}, expected)
	// require.NotEqual(t, Two{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_3(t *testing.T) {
	var expected, actual Three
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Three{}, expected)
	require.Equal(t, Three{}, actual)

	expected = Three{
		Two: Two{
			One: One{One: rando.Bool()},
			Two: rando.BoolsN(7),
		},
		Three: rando.Byte(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Three{}, expected)
	// require.NotEqual(t, Three{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_4(t *testing.T) {
	var expected, actual Four
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Four{}, expected)
	require.Equal(t, Four{}, actual)

	expected = Four{
		Three: Three{
			Two: Two{
				One: One{
					One: rando.Bool(),
				},
				Two: rando.Bools(),
			},
			Three: rando.Byte(),
		},
		Four: rando.Bytes(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Four{}, expected)
	// require.NotEqual(t, Four{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_5(t *testing.T) {
	var expected, actual Five
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Five{}, expected)
	require.Equal(t, Five{}, actual)

	expected = Five{
		Four: Four{
			Three: Three{
				Two: Two{
					One: One{
						One: rando.Bool(),
					},
					Two: rando.Bools(),
				},
				Three: rando.Byte(),
			},
			Four: rando.Bytes(),
		},
		Five: rando.Complex64(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Five{}, expected)
	// require.NotEqual(t, Five{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_6(t *testing.T) {
	var expected, actual Six
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Six{}, expected)
	require.Equal(t, Six{}, actual)

	expected = Six{
		Five: Five{
			Four: Four{
				Three: Three{
					Two: Two{
						One: One{
							One: rando.Bool(),
						},
						Two: rando.Bools(),
					},
					Three: rando.Byte(),
				},
				Four: rando.Bytes(),
			},
			Five: rando.Complex64(),
		},
		Six: rando.Complex64s(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Six{}, expected)
	// require.NotEqual(t, Six{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_7(t *testing.T) {
	var expected, actual Seven
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Seven{}, expected)
	require.Equal(t, Seven{}, actual)

	expected = Seven{
		Six: Six{
			Five: Five{
				Four: Four{
					Three: Three{
						Two: Two{
							One: One{
								One: rando.Bool(),
							},
							Two: rando.Bools(),
						},
						Three: rando.Byte(),
					},
					Four: rando.Bytes(),
				},
				Five: rando.Complex64(),
			},
			Six: rando.Complex64s(),
		},
		Seven: rando.Complex128(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Seven{}, expected)
	// require.NotEqual(t, Seven{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_8(t *testing.T) {
	var expected, actual Eight
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Eight{}, expected)
	require.Equal(t, Eight{}, actual)

	expected = Eight{
		Seven: Seven{
			Six: Six{
				Five: Five{
					Four: Four{
						Three: Three{
							Two: Two{
								One: One{
									One: rando.Bool(),
								},
								Two: rando.Bools(),
							},
							Three: rando.Byte(),
						},
						Four: rando.Bytes(),
					},
					Five: rando.Complex64(),
				},
				Six: rando.Complex64s(),
			},
			Seven: rando.Complex128(),
		},
		Eight: rando.Complex128s(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Eight{}, expected)
	// require.NotEqual(t, Eight{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_9(t *testing.T) {
	var expected, actual Nine
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Nine{}, expected)
	require.Equal(t, Nine{}, actual)

	expected = Nine{
		Eight: Eight{
			Seven: Seven{
				Six: Six{
					Five: Five{
						Four: Four{
							Three: Three{
								Two: Two{
									One: One{
										One: rando.Bool(),
									},
									Two: rando.Bools(),
								},
								Three: rando.Byte(),
							},
							Four: rando.Bytes(),
						},
						Five: rando.Complex64(),
					},
					Six: rando.Complex64s(),
				},
				Seven: rando.Complex128(),
			},
			Eight: rando.Complex128s(),
		},
		Nine: rando.Float32(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Nine{}, expected)
	// require.NotEqual(t, Nine{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_10(t *testing.T) {
	var expected, actual Ten
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Ten{}, expected)
	require.Equal(t, Ten{}, actual)

	expected = Ten{
		Nine: Nine{
			Eight: Eight{
				Seven: Seven{
					Six: Six{
						Five: Five{
							Four: Four{
								Three: Three{
									Two: Two{
										One: One{
											One: rando.Bool(),
										},
										Two: rando.Bools(),
									},
									Three: rando.Byte(),
								},
								Four: rando.Bytes(),
							},
							Five: rando.Complex64(),
						},
						Six: rando.Complex64s(),
					},
					Seven: rando.Complex128(),
				},
				Eight: rando.Complex128s(),
			},
			Nine: rando.Float32(),
		},
		Ten: rando.Float32s(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Ten{}, expected)
	// require.NotEqual(t, Ten{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_11(t *testing.T) {
	var expected, actual Eleven
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Eleven{}, expected)
	require.Equal(t, Eleven{}, actual)

	expected = Eleven{
		Ten: Ten{
			Nine: Nine{
				Eight: Eight{
					Seven: Seven{
						Six: Six{
							Five: Five{
								Four: Four{
									Three: Three{
										Two: Two{
											One: One{
												One: rando.Bool(),
											},
											Two: rando.Bools(),
										},
										Three: rando.Byte(),
									},
									Four: rando.Bytes(),
								},
								Five: rando.Complex64(),
							},
							Six: rando.Complex64s(),
						},
						Seven: rando.Complex128(),
					},
					Eight: rando.Complex128s(),
				},
				Nine: rando.Float32(),
			},
			Ten: rando.Float32s(),
		},
		Eleven: rando.Float64(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Eleven{}, expected)
	// require.NotEqual(t, Eleven{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_12(t *testing.T) {
	var expected, actual Twelve
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Twelve{}, expected)
	require.Equal(t, Twelve{}, actual)

	expected = Twelve{
		Eleven: Eleven{
			Ten: Ten{
				Nine: Nine{
					Eight: Eight{
						Seven: Seven{
							Six: Six{
								Five: Five{
									Four: Four{
										Three: Three{
											Two: Two{
												One: One{
													One: rando.Bool(),
												},
												Two: rando.Bools(),
											},
											Three: rando.Byte(),
										},
										Four: rando.Bytes(),
									},
									Five: rando.Complex64(),
								},
								Six: rando.Complex64s(),
							},
							Seven: rando.Complex128(),
						},
						Eight: rando.Complex128s(),
					},
					Nine: rando.Float32(),
				},
				Ten: rando.Float32s(),
			},
			Eleven: rando.Float64(),
		},
		Twelve: rando.Float64s(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Twelve{}, expected)
	// require.NotEqual(t, Twelve{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_13(t *testing.T) {
	var expected, actual Thirteen
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Thirteen{}, expected)
	require.Equal(t, Thirteen{}, actual)

	expected = Thirteen{
		Twelve: Twelve{
			Eleven: Eleven{
				Ten: Ten{
					Nine: Nine{
						Eight: Eight{
							Seven: Seven{
								Six: Six{
									Five: Five{
										Four: Four{
											Three: Three{
												Two: Two{
													One: One{
														One: rando.Bool(),
													},
													Two: rando.Bools(),
												},
												Three: rando.Byte(),
											},
											Four: rando.Bytes(),
										},
										Five: rando.Complex64(),
									},
									Six: rando.Complex64s(),
								},
								Seven: rando.Complex128(),
							},
							Eight: rando.Complex128s(),
						},
						Nine: rando.Float32(),
					},
					Ten: rando.Float32s(),
				},
				Eleven: rando.Float64(),
			},
			Twelve: rando.Float64s(),
		},
		Thirteen: rando.Int(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Thirteen{}, expected)
	// require.NotEqual(t, Thirteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_14(t *testing.T) {
	var expected, actual Fourteen
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Fourteen{}, expected)
	require.Equal(t, Fourteen{}, actual)

	expected = Fourteen{
		Thirteen: Thirteen{
			Twelve: Twelve{
				Eleven: Eleven{
					Ten: Ten{
						Nine: Nine{
							Eight: Eight{
								Seven: Seven{
									Six: Six{
										Five: Five{
											Four: Four{
												Three: Three{
													Two: Two{
														One: One{
															One: rando.Bool(),
														},
														Two: rando.Bools(),
													},
													Three: rando.Byte(),
												},
												Four: rando.Bytes(),
											},
											Five: rando.Complex64(),
										},
										Six: rando.Complex64s(),
									},
									Seven: rando.Complex128(),
								},
								Eight: rando.Complex128s(),
							},
							Nine: rando.Float32(),
						},
						Ten: rando.Float32s(),
					},
					Eleven: rando.Float64(),
				},
				Twelve: rando.Float64s(),
			},
			Thirteen: rando.Int(),
		},
		Fourteen: rando.Int8(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Fourteen{}, expected)
	// require.NotEqual(t, Fourteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_15(t *testing.T) {
	var expected, actual Fifteen
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Fifteen{}, expected)
	require.Equal(t, Fifteen{}, actual)

	expected = Fifteen{
		Fourteen: Fourteen{
			Thirteen: Thirteen{
				Twelve: Twelve{
					Eleven: Eleven{
						Ten: Ten{
							Nine: Nine{
								Eight: Eight{
									Seven: Seven{
										Six: Six{
											Five: Five{
												Four: Four{
													Three: Three{
														Two: Two{
															One: One{
																One: rando.Bool(),
															},
															Two: rando.Bools(),
														},
														Three: rando.Byte(),
													},
													Four: rando.Bytes(),
												},
												Five: rando.Complex64(),
											},
											Six: rando.Complex64s(),
										},
										Seven: rando.Complex128(),
									},
									Eight: rando.Complex128s(),
								},
								Nine: rando.Float32(),
							},
							Ten: rando.Float32s(),
						},
						Eleven: rando.Float64(),
					},
					Twelve: rando.Float64s(),
				},
				Thirteen: rando.Int(),
			},
			Fourteen: rando.Int8(),
		},
		Fifteen: rando.Int16(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Fifteen{}, expected)
	// require.NotEqual(t, Fifteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_16(t *testing.T) {
	var expected, actual Sixteen
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Sixteen{}, expected)
	require.Equal(t, Sixteen{}, actual)

	expected = Sixteen{
		Fifteen: Fifteen{
			Fourteen: Fourteen{
				Thirteen: Thirteen{
					Twelve: Twelve{
						Eleven: Eleven{
							Ten: Ten{
								Nine: Nine{
									Eight: Eight{
										Seven: Seven{
											Six: Six{
												Five: Five{
													Four: Four{
														Three: Three{
															Two: Two{
																One: One{
																	One: rando.Bool(),
																},
																Two: rando.Bools(),
															},
															Three: rando.Byte(),
														},
														Four: rando.Bytes(),
													},
													Five: rando.Complex64(),
												},
												Six: rando.Complex64s(),
											},
											Seven: rando.Complex128(),
										},
										Eight: rando.Complex128s(),
									},
									Nine: rando.Float32(),
								},
								Ten: rando.Float32s(),
							},
							Eleven: rando.Float64(),
						},
						Twelve: rando.Float64s(),
					},
					Thirteen: rando.Int(),
				},
				Fourteen: rando.Int8(),
			},
			Fifteen: rando.Int16(),
		},
		Sixteen: rando.Int32(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Sixteen{}, expected)
	// require.NotEqual(t, Sixteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_17(t *testing.T) {
	var expected, actual Seventeen
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Seventeen{}, expected)
	require.Equal(t, Seventeen{}, actual)

	expected = Seventeen{
		Sixteen: Sixteen{
			Fifteen: Fifteen{
				Fourteen: Fourteen{
					Thirteen: Thirteen{
						Twelve: Twelve{
							Eleven: Eleven{
								Ten: Ten{
									Nine: Nine{
										Eight: Eight{
											Seven: Seven{
												Six: Six{
													Five: Five{
														Four: Four{
															Three: Three{
																Two: Two{
																	One: One{
																		One: rando.Bool(),
																	},
																	Two: rando.Bools(),
																},
																Three: rando.Byte(),
															},
															Four: rando.Bytes(),
														},
														Five: rando.Complex64(),
													},
													Six: rando.Complex64s(),
												},
												Seven: rando.Complex128(),
											},
											Eight: rando.Complex128s(),
										},
										Nine: rando.Float32(),
									},
									Ten: rando.Float32s(),
								},
								Eleven: rando.Float64(),
							},
							Twelve: rando.Float64s(),
						},
						Thirteen: rando.Int(),
					},
					Fourteen: rando.Int8(),
				},
				Fifteen: rando.Int16(),
			},
			Sixteen: rando.Int32(),
		},
		Seventeen: rando.Int64(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Seventeen{}, expected)
	// require.NotEqual(t, Seventeen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_18(t *testing.T) {
	var expected, actual Eighteen
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Eighteen{}, expected)
	require.Equal(t, Eighteen{}, actual)

	expected = Eighteen{
		Seventeen: Seventeen{
			Sixteen: Sixteen{
				Fifteen: Fifteen{
					Fourteen: Fourteen{
						Thirteen: Thirteen{
							Twelve: Twelve{
								Eleven: Eleven{
									Ten: Ten{
										Nine: Nine{
											Eight: Eight{
												Seven: Seven{
													Six: Six{
														Five: Five{
															Four: Four{
																Three: Three{
																	Two: Two{
																		One: One{
																			One: rando.Bool(),
																		},
																		Two: rando.Bools(),
																	},
																	Three: rando.Byte(),
																},
																Four: rando.Bytes(),
															},
															Five: rando.Complex64(),
														},
														Six: rando.Complex64s(),
													},
													Seven: rando.Complex128(),
												},
												Eight: rando.Complex128s(),
											},
											Nine: rando.Float32(),
										},
										Ten: rando.Float32s(),
									},
									Eleven: rando.Float64(),
								},
								Twelve: rando.Float64s(),
							},
							Thirteen: rando.Int(),
						},
						Fourteen: rando.Int8(),
					},
					Fifteen: rando.Int16(),
				},
				Sixteen: rando.Int32(),
			},
			Seventeen: rando.Int64(),
		},
		Eighteen: rando.Ints(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Eighteen{}, expected)
	// require.NotEqual(t, Eighteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_19(t *testing.T) {
	var expected, actual Nineteen
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Nineteen{}, expected)
	require.Equal(t, Nineteen{}, actual)

	expected = Nineteen{
		Eighteen: Eighteen{
			Seventeen: Seventeen{
				Sixteen: Sixteen{
					Fifteen: Fifteen{
						Fourteen: Fourteen{
							Thirteen: Thirteen{
								Twelve: Twelve{
									Eleven: Eleven{
										Ten: Ten{
											Nine: Nine{
												Eight: Eight{
													Seven: Seven{
														Six: Six{
															Five: Five{
																Four: Four{
																	Three: Three{
																		Two: Two{
																			One: One{
																				One: rando.Bool(),
																			},
																			Two: rando.Bools(),
																		},
																		Three: rando.Byte(),
																	},
																	Four: rando.Bytes(),
																},
																Five: rando.Complex64(),
															},
															Six: rando.Complex64s(),
														},
														Seven: rando.Complex128(),
													},
													Eight: rando.Complex128s(),
												},
												Nine: rando.Float32(),
											},
											Ten: rando.Float32s(),
										},
										Eleven: rando.Float64(),
									},
									Twelve: rando.Float64s(),
								},
								Thirteen: rando.Int(),
							},
							Fourteen: rando.Int8(),
						},
						Fifteen: rando.Int16(),
					},
					Sixteen: rando.Int32(),
				},
				Seventeen: rando.Int64(),
			},
			Eighteen: rando.Ints(),
		},
		Nineteen: rando.Int8s(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Nineteen{}, expected)
	// require.NotEqual(t, Nineteen{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_20(t *testing.T) {
	var expected, actual Twenty
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Twenty{}, expected)
	require.Equal(t, Twenty{}, actual)

	expected = Twenty{
		Nineteen: Nineteen{
			Eighteen: Eighteen{
				Seventeen: Seventeen{
					Sixteen: Sixteen{
						Fifteen: Fifteen{
							Fourteen: Fourteen{
								Thirteen: Thirteen{
									Twelve: Twelve{
										Eleven: Eleven{
											Ten: Ten{
												Nine: Nine{
													Eight: Eight{
														Seven: Seven{
															Six: Six{
																Five: Five{
																	Four: Four{
																		Three: Three{
																			Two: Two{
																				One: One{
																					One: rando.Bool(),
																				},
																				Two: rando.Bools(),
																			},
																			Three: rando.Byte(),
																		},
																		Four: rando.Bytes(),
																	},
																	Five: rando.Complex64(),
																},
																Six: rando.Complex64s(),
															},
															Seven: rando.Complex128(),
														},
														Eight: rando.Complex128s(),
													},
													Nine: rando.Float32(),
												},
												Ten: rando.Float32s(),
											},
											Eleven: rando.Float64(),
										},
										Twelve: rando.Float64s(),
									},
									Thirteen: rando.Int(),
								},
								Fourteen: rando.Int8(),
							},
							Fifteen: rando.Int16(),
						},
						Sixteen: rando.Int32(),
					},
					Seventeen: rando.Int64(),
				},
				Eighteen: rando.Ints(),
			},
			Nineteen: rando.Int8s(),
		},
		Twenty: rando.Int16s(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, Twenty{}, expected)
	// require.NotEqual(t, Twenty{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_21(t *testing.T) {
	var expected, actual TwentyOne
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, TwentyOne{}, expected)
	require.Equal(t, TwentyOne{}, actual)

	expected = TwentyOne{
		Twenty: Twenty{
			Nineteen: Nineteen{
				Eighteen: Eighteen{
					Seventeen: Seventeen{
						Sixteen: Sixteen{
							Fifteen: Fifteen{
								Fourteen: Fourteen{
									Thirteen: Thirteen{
										Twelve: Twelve{
											Eleven: Eleven{
												Ten: Ten{
													Nine: Nine{
														Eight: Eight{
															Seven: Seven{
																Six: Six{
																	Five: Five{
																		Four: Four{
																			Three: Three{
																				Two: Two{
																					One: One{
																						One: rando.Bool(),
																					},
																					Two: rando.Bools(),
																				},
																				Three: rando.Byte(),
																			},
																			Four: rando.Bytes(),
																		},
																		Five: rando.Complex64(),
																	},
																	Six: rando.Complex64s(),
																},
																Seven: rando.Complex128(),
															},
															Eight: rando.Complex128s(),
														},
														Nine: rando.Float32(),
													},
													Ten: rando.Float32s(),
												},
												Eleven: rando.Float64(),
											},
											Twelve: rando.Float64s(),
										},
										Thirteen: rando.Int(),
									},
									Fourteen: rando.Int8(),
								},
								Fifteen: rando.Int16(),
							},
							Sixteen: rando.Int32(),
						},
						Seventeen: rando.Int64(),
					},
					Eighteen: rando.Ints(),
				},
				Nineteen: rando.Int8s(),
			},
			Twenty: rando.Int16s(),
		},
		TwentyOne: rando.Int32s(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, TwentyOne{}, expected)
	// require.NotEqual(t, TwentyOne{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_22(t *testing.T) {
	var expected, actual TwentyTwo
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, TwentyTwo{}, expected)
	require.Equal(t, TwentyTwo{}, actual)

	expected = TwentyTwo{
		TwentyOne: TwentyOne{
			Twenty: Twenty{
				Nineteen: Nineteen{
					Eighteen: Eighteen{
						Seventeen: Seventeen{
							Sixteen: Sixteen{
								Fifteen: Fifteen{
									Fourteen: Fourteen{
										Thirteen: Thirteen{
											Twelve: Twelve{
												Eleven: Eleven{
													Ten: Ten{
														Nine: Nine{
															Eight: Eight{
																Seven: Seven{
																	Six: Six{
																		Five: Five{
																			Four: Four{
																				Three: Three{
																					Two: Two{
																						One: One{
																							One: rando.Bool(),
																						},
																						Two: rando.Bools(),
																					},
																					Three: rando.Byte(),
																				},
																				Four: rando.Bytes(),
																			},
																			Five: rando.Complex64(),
																		},
																		Six: rando.Complex64s(),
																	},
																	Seven: rando.Complex128(),
																},
																Eight: rando.Complex128s(),
															},
															Nine: rando.Float32(),
														},
														Ten: rando.Float32s(),
													},
													Eleven: rando.Float64(),
												},
												Twelve: rando.Float64s(),
											},
											Thirteen: rando.Int(),
										},
										Fourteen: rando.Int8(),
									},
									Fifteen: rando.Int16(),
								},
								Sixteen: rando.Int32(),
							},
							Seventeen: rando.Int64(),
						},
						Eighteen: rando.Ints(),
					},
					Nineteen: rando.Int8s(),
				},
				Twenty: rando.Int16s(),
			},
			TwentyOne: rando.Int32s(),
		},
		TwentyTwo: rando.Int64s(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, TwentyTwo{}, expected)
	// require.NotEqual(t, TwentyTwo{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_23(t *testing.T) {
	var expected, actual TwentyThree
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, TwentyThree{}, expected)
	require.Equal(t, TwentyThree{}, actual)

	expected = TwentyThree{
		TwentyTwo: TwentyTwo{
			TwentyOne: TwentyOne{
				Twenty: Twenty{
					Nineteen: Nineteen{
						Eighteen: Eighteen{
							Seventeen: Seventeen{
								Sixteen: Sixteen{
									Fifteen: Fifteen{
										Fourteen: Fourteen{
											Thirteen: Thirteen{
												Twelve: Twelve{
													Eleven: Eleven{
														Ten: Ten{
															Nine: Nine{
																Eight: Eight{
																	Seven: Seven{
																		Six: Six{
																			Five: Five{
																				Four: Four{
																					Three: Three{
																						Two: Two{
																							One: One{
																								One: rando.Bool(),
																							},
																							Two: rando.Bools(),
																						},
																						Three: rando.Byte(),
																					},
																					Four: rando.Bytes(),
																				},
																				Five: rando.Complex64(),
																			},
																			Six: rando.Complex64s(),
																		},
																		Seven: rando.Complex128(),
																	},
																	Eight: rando.Complex128s(),
																},
																Nine: rando.Float32(),
															},
															Ten: rando.Float32s(),
														},
														Eleven: rando.Float64(),
													},
													Twelve: rando.Float64s(),
												},
												Thirteen: rando.Int(),
											},
											Fourteen: rando.Int8(),
										},
										Fifteen: rando.Int16(),
									},
									Sixteen: rando.Int32(),
								},
								Seventeen: rando.Int64(),
							},
							Eighteen: rando.Ints(),
						},
						Nineteen: rando.Int8s(),
					},
					Twenty: rando.Int16s(),
				},
				TwentyOne: rando.Int32s(),
			},
			TwentyTwo: rando.Int64s(),
		},
		TwentyThree: rando.Uint(),
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	// require.NotEqual(t, TwentyThree{}, expected)
	// require.NotEqual(t, TwentyThree{}, actual)
	require.Equal(t, expected, actual)
}
