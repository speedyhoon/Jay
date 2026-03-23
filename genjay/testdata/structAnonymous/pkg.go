package main

import "time"

type One struct {
	One bool
}

type Two struct {
	One struct {
		One bool
	}
	Two []bool
}

type Three struct {
	Two struct {
		One struct {
			One bool
		}
		Two []bool
	}
	Three byte
}

type Four struct {
	Three struct {
		Two struct {
			One struct {
				One bool
			}
			Two []bool
		}
		Three byte
	}
	Four []byte
}

type Five struct {
	Four struct {
		Three struct {
			Two struct {
				One struct {
					One bool
				}
				Two []bool
			}
			Three byte
		}
		Four []byte
	}
	Five complex64
}

type Six struct {
	Five struct {
		Four struct {
			Three struct {
				Two struct {
					One struct {
						One bool
					}
					Two []bool
				}
				Three byte
			}
			Four []byte
		}
		Five complex64
	}
	Six []complex64
}

type Seven struct {
	Six struct {
		Five struct {
			Four struct {
				Three struct {
					Two struct {
						One struct {
							One bool
						}
						Two []bool
					}
					Three byte
				}
				Four []byte
			}
			Five complex64
		}
		Six []complex64
	}
	Seven complex128
}

type Eight struct {
	Seven struct {
		Six struct {
			Five struct {
				Four struct {
					Three struct {
						Two struct {
							One struct {
								One bool
							}
							Two []bool
						}
						Three byte
					}
					Four []byte
				}
				Five complex64
			}
			Six []complex64
		}
		Seven complex128
	}
	Eight []complex128
}

type Nine struct {
	Eight struct {
		Seven struct {
			Six struct {
				Five struct {
					Four struct {
						Three struct {
							Two struct {
								One struct {
									One bool
								}
								Two []bool
							}
							Three byte
						}
						Four []byte
					}
					Five complex64
				}
				Six []complex64
			}
			Seven complex128
		}
		Eight []complex128
	}
	Nine float32
}

type Ten struct {
	Nine struct {
		Eight struct {
			Seven struct {
				Six struct {
					Five struct {
						Four struct {
							Three struct {
								Two struct {
									One struct {
										One bool
									}
									Two []bool
								}
								Three byte
							}
							Four []byte
						}
						Five complex64
					}
					Six []complex64
				}
				Seven complex128
			}
			Eight []complex128
		}
		Nine float32
	}
	Ten []float32
}

type Eleven struct {
	Ten struct {
		Nine struct {
			Eight struct {
				Seven struct {
					Six struct {
						Five struct {
							Four struct {
								Three struct {
									Two struct {
										One struct {
											One bool
										}
										Two []bool
									}
									Three byte
								}
								Four []byte
							}
							Five complex64
						}
						Six []complex64
					}
					Seven complex128
				}
				Eight []complex128
			}
			Nine float32
		}
		Ten []float32
	}
	Eleven float64
}

type Twelve struct {
	Eleven struct {
		Ten struct {
			Nine struct {
				Eight struct {
					Seven struct {
						Six struct {
							Five struct {
								Four struct {
									Three struct {
										Two struct {
											One struct {
												One bool
											}
											Two []bool
										}
										Three byte
									}
									Four []byte
								}
								Five complex64
							}
							Six []complex64
						}
						Seven complex128
					}
					Eight []complex128
				}
				Nine float32
			}
			Ten []float32
		}
		Eleven float64
	}
	Twelve []float64
}

type Thirteen struct {
	Twelve struct {
		Eleven struct {
			Ten struct {
				Nine struct {
					Eight struct {
						Seven struct {
							Six struct {
								Five struct {
									Four struct {
										Three struct {
											Two struct {
												One struct {
													One bool
												}
												Two []bool
											}
											Three byte
										}
										Four []byte
									}
									Five complex64
								}
								Six []complex64
							}
							Seven complex128
						}
						Eight []complex128
					}
					Nine float32
				}
				Ten []float32
			}
			Eleven float64
		}
		Twelve []float64
	}
	Thirteen int
}

type Fourteen struct {
	Thirteen struct {
		Twelve struct {
			Eleven struct {
				Ten struct {
					Nine struct {
						Eight struct {
							Seven struct {
								Six struct {
									Five struct {
										Four struct {
											Three struct {
												Two struct {
													One struct {
														One bool
													}
													Two []bool
												}
												Three byte
											}
											Four []byte
										}
										Five complex64
									}
									Six []complex64
								}
								Seven complex128
							}
							Eight []complex128
						}
						Nine float32
					}
					Ten []float32
				}
				Eleven float64
			}
			Twelve []float64
		}
		Thirteen int
	}
	Fourteen int8
}

type Fifteen struct {
	Fourteen struct {
		Thirteen struct {
			Twelve struct {
				Eleven struct {
					Ten struct {
						Nine struct {
							Eight struct {
								Seven struct {
									Six struct {
										Five struct {
											Four struct {
												Three struct {
													Two struct {
														One struct {
															One bool
														}
														Two []bool
													}
													Three byte
												}
												Four []byte
											}
											Five complex64
										}
										Six []complex64
									}
									Seven complex128
								}
								Eight []complex128
							}
							Nine float32
						}
						Ten []float32
					}
					Eleven float64
				}
				Twelve []float64
			}
			Thirteen int
		}
		Fourteen int8
	}
	Fifteen int16
}

type Sixteen struct {
	Fifteen struct {
		Fourteen struct {
			Thirteen struct {
				Twelve struct {
					Eleven struct {
						Ten struct {
							Nine struct {
								Eight struct {
									Seven struct {
										Six struct {
											Five struct {
												Four struct {
													Three struct {
														Two struct {
															One struct {
																One bool
															}
															Two []bool
														}
														Three byte
													}
													Four []byte
												}
												Five complex64
											}
											Six []complex64
										}
										Seven complex128
									}
									Eight []complex128
								}
								Nine float32
							}
							Ten []float32
						}
						Eleven float64
					}
					Twelve []float64
				}
				Thirteen int
			}
			Fourteen int8
		}
		Fifteen int16
	}
	Sixteen int32
}

type Seventeen struct {
	Sixteen struct {
		Fifteen struct {
			Fourteen struct {
				Thirteen struct {
					Twelve struct {
						Eleven struct {
							Ten struct {
								Nine struct {
									Eight struct {
										Seven struct {
											Six struct {
												Five struct {
													Four struct {
														Three struct {
															Two struct {
																One struct {
																	One bool
																}
																Two []bool
															}
															Three byte
														}
														Four []byte
													}
													Five complex64
												}
												Six []complex64
											}
											Seven complex128
										}
										Eight []complex128
									}
									Nine float32
								}
								Ten []float32
							}
							Eleven float64
						}
						Twelve []float64
					}
					Thirteen int
				}
				Fourteen int8
			}
			Fifteen int16
		}
		Sixteen int32
	}
	Seventeen int64
}

type Eighteen struct {
	Seventeen struct {
		Sixteen struct {
			Fifteen struct {
				Fourteen struct {
					Thirteen struct {
						Twelve struct {
							Eleven struct {
								Ten struct {
									Nine struct {
										Eight struct {
											Seven struct {
												Six struct {
													Five struct {
														Four struct {
															Three struct {
																Two struct {
																	One struct {
																		One bool
																	}
																	Two []bool
																}
																Three byte
															}
															Four []byte
														}
														Five complex64
													}
													Six []complex64
												}
												Seven complex128
											}
											Eight []complex128
										}
										Nine float32
									}
									Ten []float32
								}
								Eleven float64
							}
							Twelve []float64
						}
						Thirteen int
					}
					Fourteen int8
				}
				Fifteen int16
			}
			Sixteen int32
		}
		Seventeen int64
	}
	Eighteen []int
}

type Nineteen struct {
	Eighteen struct {
		Seventeen struct {
			Sixteen struct {
				Fifteen struct {
					Fourteen struct {
						Thirteen struct {
							Twelve struct {
								Eleven struct {
									Ten struct {
										Nine struct {
											Eight struct {
												Seven struct {
													Six struct {
														Five struct {
															Four struct {
																Three struct {
																	Two struct {
																		One struct {
																			One bool
																		}
																		Two []bool
																	}
																	Three byte
																}
																Four []byte
															}
															Five complex64
														}
														Six []complex64
													}
													Seven complex128
												}
												Eight []complex128
											}
											Nine float32
										}
										Ten []float32
									}
									Eleven float64
								}
								Twelve []float64
							}
							Thirteen int
						}
						Fourteen int8
					}
					Fifteen int16
				}
				Sixteen int32
			}
			Seventeen int64
		}
		Eighteen []int
	}
	Nineteen []int8
}

type Twenty struct {
	Nineteen struct {
		Eighteen struct {
			Seventeen struct {
				Sixteen struct {
					Fifteen struct {
						Fourteen struct {
							Thirteen struct {
								Twelve struct {
									Eleven struct {
										Ten struct {
											Nine struct {
												Eight struct {
													Seven struct {
														Six struct {
															Five struct {
																Four struct {
																	Three struct {
																		Two struct {
																			One struct {
																				One bool
																			}
																			Two []bool
																		}
																		Three byte
																	}
																	Four []byte
																}
																Five complex64
															}
															Six []complex64
														}
														Seven complex128
													}
													Eight []complex128
												}
												Nine float32
											}
											Ten []float32
										}
										Eleven float64
									}
									Twelve []float64
								}
								Thirteen int
							}
							Fourteen int8
						}
						Fifteen int16
					}
					Sixteen int32
				}
				Seventeen int64
			}
			Eighteen []int
		}
		Nineteen []int8
	}
	Twenty []int16
}

type TwentyOne struct {
	Twenty struct {
		Nineteen struct {
			Eighteen struct {
				Seventeen struct {
					Sixteen struct {
						Fifteen struct {
							Fourteen struct {
								Thirteen struct {
									Twelve struct {
										Eleven struct {
											Ten struct {
												Nine struct {
													Eight struct {
														Seven struct {
															Six struct {
																Five struct {
																	Four struct {
																		Three struct {
																			Two struct {
																				One struct {
																					One bool
																				}
																				Two []bool
																			}
																			Three byte
																		}
																		Four []byte
																	}
																	Five complex64
																}
																Six []complex64
															}
															Seven complex128
														}
														Eight []complex128
													}
													Nine float32
												}
												Ten []float32
											}
											Eleven float64
										}
										Twelve []float64
									}
									Thirteen int
								}
								Fourteen int8
							}
							Fifteen int16
						}
						Sixteen int32
					}
					Seventeen int64
				}
				Eighteen []int
			}
			Nineteen []int8
		}
		Twenty []int16
	}
	TwentyOne []int32
}

type TwentyTwo struct {
	TwentyOne struct {
		Twenty struct {
			Nineteen struct {
				Eighteen struct {
					Seventeen struct {
						Sixteen struct {
							Fifteen struct {
								Fourteen struct {
									Thirteen struct {
										Twelve struct {
											Eleven struct {
												Ten struct {
													Nine struct {
														Eight struct {
															Seven struct {
																Six struct {
																	Five struct {
																		Four struct {
																			Three struct {
																				Two struct {
																					One struct {
																						One bool
																					}
																					Two []bool
																				}
																				Three byte
																			}
																			Four []byte
																		}
																		Five complex64
																	}
																	Six []complex64
																}
																Seven complex128
															}
															Eight []complex128
														}
														Nine float32
													}
													Ten []float32
												}
												Eleven float64
											}
											Twelve []float64
										}
										Thirteen int
									}
									Fourteen int8
								}
								Fifteen int16
							}
							Sixteen int32
						}
						Seventeen int64
					}
					Eighteen []int
				}
				Nineteen []int8
			}
			Twenty []int16
		}
		TwentyOne []int32
	}
	TwentyTwo []int64
}

type TwentyThree struct {
	TwentyTwo struct {
		TwentyOne struct {
			Twenty struct {
				Nineteen struct {
					Eighteen struct {
						Seventeen struct {
							Sixteen struct {
								Fifteen struct {
									Fourteen struct {
										Thirteen struct {
											Twelve struct {
												Eleven struct {
													Ten struct {
														Nine struct {
															Eight struct {
																Seven struct {
																	Six struct {
																		Five struct {
																			Four struct {
																				Three struct {
																					Two struct {
																						One struct {
																							One bool
																						}
																						Two []bool
																					}
																					Three byte
																				}
																				Four []byte
																			}
																			Five complex64
																		}
																		Six []complex64
																	}
																	Seven complex128
																}
																Eight []complex128
															}
															Nine float32
														}
														Ten []float32
													}
													Eleven float64
												}
												Twelve []float64
											}
											Thirteen int
										}
										Fourteen int8
									}
									Fifteen int16
								}
								Sixteen int32
							}
							Seventeen int64
						}
						Eighteen []int
					}
					Nineteen []int8
				}
				Twenty []int16
			}
			TwentyOne []int32
		}
		TwentyTwo []int64
	}
	TwentyThree uint
}

type TwentyFour struct {
	TwentyThree struct {
		TwentyTwo struct {
			TwentyOne struct {
				Twenty struct {
					Nineteen struct {
						Eighteen struct {
							Seventeen struct {
								Sixteen struct {
									Fifteen struct {
										Fourteen struct {
											Thirteen struct {
												Twelve struct {
													Eleven struct {
														Ten struct {
															Nine struct {
																Eight struct {
																	Seven struct {
																		Six struct {
																			Five struct {
																				Four struct {
																					Three struct {
																						Two struct {
																							One struct {
																								One bool
																							}
																							Two []bool
																						}
																						Three byte
																					}
																					Four []byte
																				}
																				Five complex64
																			}
																			Six []complex64
																		}
																		Seven complex128
																	}
																	Eight []complex128
																}
																Nine float32
															}
															Ten []float32
														}
														Eleven float64
													}
													Twelve []float64
												}
												Thirteen int
											}
											Fourteen int8
										}
										Fifteen int16
									}
									Sixteen int32
								}
								Seventeen int64
							}
							Eighteen []int
						}
						Nineteen []int8
					}
					Twenty []int16
				}
				TwentyOne []int32
			}
			TwentyTwo []int64
		}
		TwentyThree uint
	}
	TwentyFour uint16
}

type TwentyFive struct {
	TwentyFour struct {
		TwentyThree struct {
			TwentyTwo struct {
				TwentyOne struct {
					Twenty struct {
						Nineteen struct {
							Eighteen struct {
								Seventeen struct {
									Sixteen struct {
										Fifteen struct {
											Fourteen struct {
												Thirteen struct {
													Twelve struct {
														Eleven struct {
															Ten struct {
																Nine struct {
																	Eight struct {
																		Seven struct {
																			Six struct {
																				Five struct {
																					Four struct {
																						Three struct {
																							Two struct {
																								One struct {
																									One bool
																								}
																								Two []bool
																							}
																							Three byte
																						}
																						Four []byte
																					}
																					Five complex64
																				}
																				Six []complex64
																			}
																			Seven complex128
																		}
																		Eight []complex128
																	}
																	Nine float32
																}
																Ten []float32
															}
															Eleven float64
														}
														Twelve []float64
													}
													Thirteen int
												}
												Fourteen int8
											}
											Fifteen int16
										}
										Sixteen int32
									}
									Seventeen int64
								}
								Eighteen []int
							}
							Nineteen []int8
						}
						Twenty []int16
					}
					TwentyOne []int32
				}
				TwentyTwo []int64
			}
			TwentyThree uint
		}
		TwentyFour uint16
	}
	TwentyFive uint32
}

type TwentySix struct {
	TwentyFive struct {
		TwentyFour struct {
			TwentyThree struct {
				TwentyTwo struct {
					TwentyOne struct {
						Twenty struct {
							Nineteen struct {
								Eighteen struct {
									Seventeen struct {
										Sixteen struct {
											Fifteen struct {
												Fourteen struct {
													Thirteen struct {
														Twelve struct {
															Eleven struct {
																Ten struct {
																	Nine struct {
																		Eight struct {
																			Seven struct {
																				Six struct {
																					Five struct {
																						Four struct {
																							Three struct {
																								Two struct {
																									One struct {
																										One bool
																									}
																									Two []bool
																								}
																								Three byte
																							}
																							Four []byte
																						}
																						Five complex64
																					}
																					Six []complex64
																				}
																				Seven complex128
																			}
																			Eight []complex128
																		}
																		Nine float32
																	}
																	Ten []float32
																}
																Eleven float64
															}
															Twelve []float64
														}
														Thirteen int
													}
													Fourteen int8
												}
												Fifteen int16
											}
											Sixteen int32
										}
										Seventeen int64
									}
									Eighteen []int
								}
								Nineteen []int8
							}
							Twenty []int16
						}
						TwentyOne []int32
					}
					TwentyTwo []int64
				}
				TwentyThree uint
			}
			TwentyFour uint16
		}
		TwentyFive uint32
	}
	TwentySix uint64
}

type TwentySeven struct {
	TwentySix struct {
		TwentyFive struct {
			TwentyFour struct {
				TwentyThree struct {
					TwentyTwo struct {
						TwentyOne struct {
							Twenty struct {
								Nineteen struct {
									Eighteen struct {
										Seventeen struct {
											Sixteen struct {
												Fifteen struct {
													Fourteen struct {
														Thirteen struct {
															Twelve struct {
																Eleven struct {
																	Ten struct {
																		Nine struct {
																			Eight struct {
																				Seven struct {
																					Six struct {
																						Five struct {
																							Four struct {
																								Three struct {
																									Two struct {
																										One struct {
																											One bool
																										}
																										Two []bool
																									}
																									Three byte
																								}
																								Four []byte
																							}
																							Five complex64
																						}
																						Six []complex64
																					}
																					Seven complex128
																				}
																				Eight []complex128
																			}
																			Nine float32
																		}
																		Ten []float32
																	}
																	Eleven float64
																}
																Twelve []float64
															}
															Thirteen int
														}
														Fourteen int8
													}
													Fifteen int16
												}
												Sixteen int32
											}
											Seventeen int64
										}
										Eighteen []int
									}
									Nineteen []int8
								}
								Twenty []int16
							}
							TwentyOne []int32
						}
						TwentyTwo []int64
					}
					TwentyThree uint
				}
				TwentyFour uint16
			}
			TwentyFive uint32
		}
		TwentySix uint64
	}
	TwentySeven []uint
}

type TwentyEight struct {
	TwentySeven struct {
		TwentySix struct {
			TwentyFive struct {
				TwentyFour struct {
					TwentyThree struct {
						TwentyTwo struct {
							TwentyOne struct {
								Twenty struct {
									Nineteen struct {
										Eighteen struct {
											Seventeen struct {
												Sixteen struct {
													Fifteen struct {
														Fourteen struct {
															Thirteen struct {
																Twelve struct {
																	Eleven struct {
																		Ten struct {
																			Nine struct {
																				Eight struct {
																					Seven struct {
																						Six struct {
																							Five struct {
																								Four struct {
																									Three struct {
																										Two struct {
																											One struct {
																												One bool
																											}
																											Two []bool
																										}
																										Three byte
																									}
																									Four []byte
																								}
																								Five complex64
																							}
																							Six []complex64
																						}
																						Seven complex128
																					}
																					Eight []complex128
																				}
																				Nine float32
																			}
																			Ten []float32
																		}
																		Eleven float64
																	}
																	Twelve []float64
																}
																Thirteen int
															}
															Fourteen int8
														}
														Fifteen int16
													}
													Sixteen int32
												}
												Seventeen int64
											}
											Eighteen []int
										}
										Nineteen []int8
									}
									Twenty []int16
								}
								TwentyOne []int32
							}
							TwentyTwo []int64
						}
						TwentyThree uint
					}
					TwentyFour uint16
				}
				TwentyFive uint32
			}
			TwentySix uint64
		}
		TwentySeven []uint
	}
	TwentyEight []uint16
}

type TwentyNine struct {
	TwentyEight struct {
		TwentySeven struct {
			TwentySix struct {
				TwentyFive struct {
					TwentyFour struct {
						TwentyThree struct {
							TwentyTwo struct {
								TwentyOne struct {
									Twenty struct {
										Nineteen struct {
											Eighteen struct {
												Seventeen struct {
													Sixteen struct {
														Fifteen struct {
															Fourteen struct {
																Thirteen struct {
																	Twelve struct {
																		Eleven struct {
																			Ten struct {
																				Nine struct {
																					Eight struct {
																						Seven struct {
																							Six struct {
																								Five struct {
																									Four struct {
																										Three struct {
																											Two struct {
																												One struct {
																													One bool
																												}
																												Two []bool
																											}
																											Three byte
																										}
																										Four []byte
																									}
																									Five complex64
																								}
																								Six []complex64
																							}
																							Seven complex128
																						}
																						Eight []complex128
																					}
																					Nine float32
																				}
																				Ten []float32
																			}
																			Eleven float64
																		}
																		Twelve []float64
																	}
																	Thirteen int
																}
																Fourteen int8
															}
															Fifteen int16
														}
														Sixteen int32
													}
													Seventeen int64
												}
												Eighteen []int
											}
											Nineteen []int8
										}
										Twenty []int16
									}
									TwentyOne []int32
								}
								TwentyTwo []int64
							}
							TwentyThree uint
						}
						TwentyFour uint16
					}
					TwentyFive uint32
				}
				TwentySix uint64
			}
			TwentySeven []uint
		}
		TwentyEight []uint16
	}
	TwentyNine []uint32
}

type Thirty struct {
	TwentyNine struct {
		TwentyEight struct {
			TwentySeven struct {
				TwentySix struct {
					TwentyFive struct {
						TwentyFour struct {
							TwentyThree struct {
								TwentyTwo struct {
									TwentyOne struct {
										Twenty struct {
											Nineteen struct {
												Eighteen struct {
													Seventeen struct {
														Sixteen struct {
															Fifteen struct {
																Fourteen struct {
																	Thirteen struct {
																		Twelve struct {
																			Eleven struct {
																				Ten struct {
																					Nine struct {
																						Eight struct {
																							Seven struct {
																								Six struct {
																									Five struct {
																										Four struct {
																											Three struct {
																												Two struct {
																													One struct {
																														One bool
																													}
																													Two []bool
																												}
																												Three byte
																											}
																											Four []byte
																										}
																										Five complex64
																									}
																									Six []complex64
																								}
																								Seven complex128
																							}
																							Eight []complex128
																						}
																						Nine float32
																					}
																					Ten []float32
																				}
																				Eleven float64
																			}
																			Twelve []float64
																		}
																		Thirteen int
																	}
																	Fourteen int8
																}
																Fifteen int16
															}
															Sixteen int32
														}
														Seventeen int64
													}
													Eighteen []int
												}
												Nineteen []int8
											}
											Twenty []int16
										}
										TwentyOne []int32
									}
									TwentyTwo []int64
								}
								TwentyThree uint
							}
							TwentyFour uint16
						}
						TwentyFive uint32
					}
					TwentySix uint64
				}
				TwentySeven []uint
			}
			TwentyEight []uint16
		}
		TwentyNine []uint32
	}
	Thirty []uint64
}

type ThirtyOne struct {
	Thirty struct {
		TwentyNine struct {
			TwentyEight struct {
				TwentySeven struct {
					TwentySix struct {
						TwentyFive struct {
							TwentyFour struct {
								TwentyThree struct {
									TwentyTwo struct {
										TwentyOne struct {
											Twenty struct {
												Nineteen struct {
													Eighteen struct {
														Seventeen struct {
															Sixteen struct {
																Fifteen struct {
																	Fourteen struct {
																		Thirteen struct {
																			Twelve struct {
																				Eleven struct {
																					Ten struct {
																						Nine struct {
																							Eight struct {
																								Seven struct {
																									Six struct {
																										Five struct {
																											Four struct {
																												Three struct {
																													Two struct {
																														One struct {
																															One bool
																														}
																														Two []bool
																													}
																													Three byte
																												}
																												Four []byte
																											}
																											Five complex64
																										}
																										Six []complex64
																									}
																									Seven complex128
																								}
																								Eight []complex128
																							}
																							Nine float32
																						}
																						Ten []float32
																					}
																					Eleven float64
																				}
																				Twelve []float64
																			}
																			Thirteen int
																		}
																		Fourteen int8
																	}
																	Fifteen int16
																}
																Sixteen int32
															}
															Seventeen int64
														}
														Eighteen []int
													}
													Nineteen []int8
												}
												Twenty []int16
											}
											TwentyOne []int32
										}
										TwentyTwo []int64
									}
									TwentyThree uint
								}
								TwentyFour uint16
							}
							TwentyFive uint32
						}
						TwentySix uint64
					}
					TwentySeven []uint
				}
				TwentyEight []uint16
			}
			TwentyNine []uint32
		}
		Thirty []uint64
	}
	ThirtyOne string
}

type ThirtyTwo struct {
	ThirtyOne struct {
		Thirty struct {
			TwentyNine struct {
				TwentyEight struct {
					TwentySeven struct {
						TwentySix struct {
							TwentyFive struct {
								TwentyFour struct {
									TwentyThree struct {
										TwentyTwo struct {
											TwentyOne struct {
												Twenty struct {
													Nineteen struct {
														Eighteen struct {
															Seventeen struct {
																Sixteen struct {
																	Fifteen struct {
																		Fourteen struct {
																			Thirteen struct {
																				Twelve struct {
																					Eleven struct {
																						Ten struct {
																							Nine struct {
																								Eight struct {
																									Seven struct {
																										Six struct {
																											Five struct {
																												Four struct {
																													Three struct {
																														Two struct {
																															One struct {
																																One bool
																															}
																															Two []bool
																														}
																														Three byte
																													}
																													Four []byte
																												}
																												Five complex64
																											}
																											Six []complex64
																										}
																										Seven complex128
																									}
																									Eight []complex128
																								}
																								Nine float32
																							}
																							Ten []float32
																						}
																						Eleven float64
																					}
																					Twelve []float64
																				}
																				Thirteen int
																			}
																			Fourteen int8
																		}
																		Fifteen int16
																	}
																	Sixteen int32
																}
																Seventeen int64
															}
															Eighteen []int
														}
														Nineteen []int8
													}
													Twenty []int16
												}
												TwentyOne []int32
											}
											TwentyTwo []int64
										}
										TwentyThree uint
									}
									TwentyFour uint16
								}
								TwentyFive uint32
							}
							TwentySix uint64
						}
						TwentySeven []uint
					}
					TwentyEight []uint16
				}
				TwentyNine []uint32
			}
			Thirty []uint64
		}
		ThirtyOne string
	}
	ThirtyTwo []string
}

type ThirtyThree struct {
	ThirtyTwo struct {
		ThirtyOne struct {
			Thirty struct {
				TwentyNine struct {
					TwentyEight struct {
						TwentySeven struct {
							TwentySix struct {
								TwentyFive struct {
									TwentyFour struct {
										TwentyThree struct {
											TwentyTwo struct {
												TwentyOne struct {
													Twenty struct {
														Nineteen struct {
															Eighteen struct {
																Seventeen struct {
																	Sixteen struct {
																		Fifteen struct {
																			Fourteen struct {
																				Thirteen struct {
																					Twelve struct {
																						Eleven struct {
																							Ten struct {
																								Nine struct {
																									Eight struct {
																										Seven struct {
																											Six struct {
																												Five struct {
																													Four struct {
																														Three struct {
																															Two struct {
																																One struct {
																																	One bool
																																}
																																Two []bool
																															}
																															Three byte
																														}
																														Four []byte
																													}
																													Five complex64
																												}
																												Six []complex64
																											}
																											Seven complex128
																										}
																										Eight []complex128
																									}
																									Nine float32
																								}
																								Ten []float32
																							}
																							Eleven float64
																						}
																						Twelve []float64
																					}
																					Thirteen int
																				}
																				Fourteen int8
																			}
																			Fifteen int16
																		}
																		Sixteen int32
																	}
																	Seventeen int64
																}
																Eighteen []int
															}
															Nineteen []int8
														}
														Twenty []int16
													}
													TwentyOne []int32
												}
												TwentyTwo []int64
											}
											TwentyThree uint
										}
										TwentyFour uint16
									}
									TwentyFive uint32
								}
								TwentySix uint64
							}
							TwentySeven []uint
						}
						TwentyEight []uint16
					}
					TwentyNine []uint32
				}
				Thirty []uint64
			}
			ThirtyOne string
		}
		ThirtyTwo []string
	}
	ThirtyThree time.Time
}

type ThirtyFour struct {
	ThirtyThree struct {
		ThirtyTwo struct {
			ThirtyOne struct {
				Thirty struct {
					TwentyNine struct {
						TwentyEight struct {
							TwentySeven struct {
								TwentySix struct {
									TwentyFive struct {
										TwentyFour struct {
											TwentyThree struct {
												TwentyTwo struct {
													TwentyOne struct {
														Twenty struct {
															Nineteen struct {
																Eighteen struct {
																	Seventeen struct {
																		Sixteen struct {
																			Fifteen struct {
																				Fourteen struct {
																					Thirteen struct {
																						Twelve struct {
																							Eleven struct {
																								Ten struct {
																									Nine struct {
																										Eight struct {
																											Seven struct {
																												Six struct {
																													Five struct {
																														Four struct {
																															Three struct {
																																Two struct {
																																	One struct {
																																		One bool
																																	}
																																	Two []bool
																																}
																																Three byte
																															}
																															Four []byte
																														}
																														Five complex64
																													}
																													Six []complex64
																												}
																												Seven complex128
																											}
																											Eight []complex128
																										}
																										Nine float32
																									}
																									Ten []float32
																								}
																								Eleven float64
																							}
																							Twelve []float64
																						}
																						Thirteen int
																					}
																					Fourteen int8
																				}
																				Fifteen int16
																			}
																			Sixteen int32
																		}
																		Seventeen int64
																	}
																	Eighteen []int
																}
																Nineteen []int8
															}
															Twenty []int16
														}
														TwentyOne []int32
													}
													TwentyTwo []int64
												}
												TwentyThree uint
											}
											TwentyFour uint16
										}
										TwentyFive uint32
									}
									TwentySix uint64
								}
								TwentySeven []uint
							}
							TwentyEight []uint16
						}
						TwentyNine []uint32
					}
					Thirty []uint64
				}
				ThirtyOne string
			}
			ThirtyTwo []string
		}
		ThirtyThree time.Time
	}
	ThirtyFour []time.Time
}

type ThirtyFive struct {
	ThirtyFour struct {
		ThirtyThree struct {
			ThirtyTwo struct {
				ThirtyOne struct {
					Thirty struct {
						TwentyNine struct {
							TwentyEight struct {
								TwentySeven struct {
									TwentySix struct {
										TwentyFive struct {
											TwentyFour struct {
												TwentyThree struct {
													TwentyTwo struct {
														TwentyOne struct {
															Twenty struct {
																Nineteen struct {
																	Eighteen struct {
																		Seventeen struct {
																			Sixteen struct {
																				Fifteen struct {
																					Fourteen struct {
																						Thirteen struct {
																							Twelve struct {
																								Eleven struct {
																									Ten struct {
																										Nine struct {
																											Eight struct {
																												Seven struct {
																													Six struct {
																														Five struct {
																															Four struct {
																																Three struct {
																																	Two struct {
																																		One struct {
																																			One bool
																																		}
																																		Two []bool
																																	}
																																	Three byte
																																}
																																Four []byte
																															}
																															Five complex64
																														}
																														Six []complex64
																													}
																													Seven complex128
																												}
																												Eight []complex128
																											}
																											Nine float32
																										}
																										Ten []float32
																									}
																									Eleven float64
																								}
																								Twelve []float64
																							}
																							Thirteen int
																						}
																						Fourteen int8
																					}
																					Fifteen int16
																				}
																				Sixteen int32
																			}
																			Seventeen int64
																		}
																		Eighteen []int
																	}
																	Nineteen []int8
																}
																Twenty []int16
															}
															TwentyOne []int32
														}
														TwentyTwo []int64
													}
													TwentyThree uint
												}
												TwentyFour uint16
											}
											TwentyFive uint32
										}
										TwentySix uint64
									}
									TwentySeven []uint
								}
								TwentyEight []uint16
							}
							TwentyNine []uint32
						}
						Thirty []uint64
					}
					ThirtyOne string
				}
				ThirtyTwo []string
			}
			ThirtyThree time.Time
		}
		ThirtyFour []time.Time
	}
	ThirtyFive time.Duration
}

type ThirtySix struct {
	ThirtyFive struct {
		ThirtyFour struct {
			ThirtyThree struct {
				ThirtyTwo struct {
					ThirtyOne struct {
						Thirty struct {
							TwentyNine struct {
								TwentyEight struct {
									TwentySeven struct {
										TwentySix struct {
											TwentyFive struct {
												TwentyFour struct {
													TwentyThree struct {
														TwentyTwo struct {
															TwentyOne struct {
																Twenty struct {
																	Nineteen struct {
																		Eighteen struct {
																			Seventeen struct {
																				Sixteen struct {
																					Fifteen struct {
																						Fourteen struct {
																							Thirteen struct {
																								Twelve struct {
																									Eleven struct {
																										Ten struct {
																											Nine struct {
																												Eight struct {
																													Seven struct {
																														Six struct {
																															Five struct {
																																Four struct {
																																	Three struct {
																																		Two struct {
																																			One struct {
																																				One bool
																																			}
																																			Two []bool
																																		}
																																		Three byte
																																	}
																																	Four []byte
																																}
																																Five complex64
																															}
																															Six []complex64
																														}
																														Seven complex128
																													}
																													Eight []complex128
																												}
																												Nine float32
																											}
																											Ten []float32
																										}
																										Eleven float64
																									}
																									Twelve []float64
																								}
																								Thirteen int
																							}
																							Fourteen int8
																						}
																						Fifteen int16
																					}
																					Sixteen int32
																				}
																				Seventeen int64
																			}
																			Eighteen []int
																		}
																		Nineteen []int8
																	}
																	Twenty []int16
																}
																TwentyOne []int32
															}
															TwentyTwo []int64
														}
														TwentyThree uint
													}
													TwentyFour uint16
												}
												TwentyFive uint32
											}
											TwentySix uint64
										}
										TwentySeven []uint
									}
									TwentyEight []uint16
								}
								TwentyNine []uint32
							}
							Thirty []uint64
						}
						ThirtyOne string
					}
					ThirtyTwo []string
				}
				ThirtyThree time.Time
			}
			ThirtyFour []time.Time
		}
		ThirtyFive time.Duration
	}
	ThirtySix []time.Duration
}
