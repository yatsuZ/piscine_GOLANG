package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n == 1 {
		PrintCombUN()
	} else if n == 2 {
		PrintCombDeux()
	} else if n == 3 {
		PrintCombTrois()
	} else if n == 4 {
		PrintCombQuatre()
	} else if n == 5 {
		PrintCombCinq()
	} else if n == 6 {
		PrintCombSix()
	} else if n == 7 {
		PrintCombSept()
	} else if n == 8 {
		PrintCombHuit()
	} else if n == 9 {
		PrintCombNeuf()
	}
}

func PrintCombNeuf() {
	for BOSS := '0'; BOSS <= '9'; BOSS++ {
		for Mm := '0'; Mm <= '9'; Mm++ {
			for Mc := '0'; Mc <= '9'; Mc++ {
				for Md := '0'; Md <= '9'; Md++ {
					for Mu := '0'; Mu <= '9'; Mu++ {
						for m := '0'; m <= '9'; m++ {
							for c := '0'; c <= '9'; c++ {
								for d := '0'; d <= '9'; d++ {
									for u := '0'; u <= '9'; u++ {
										if BOSS < Mm && Mm < Mc && Mc < Md && Md < Mu && Mu < m && m < c && c < d && d < u {
											z01.PrintRune(BOSS)
											z01.PrintRune(Mm)
											z01.PrintRune(Mc)
											z01.PrintRune(Md)
											z01.PrintRune(Mu)
											z01.PrintRune(m)
											z01.PrintRune(c)
											z01.PrintRune(d)
											z01.PrintRune(u)
											if BOSS == '1' && Mm == '2' && Mc == '3' && Md == '4' && Mu == '5' && m == '6' && c == '7' && d == '8' && u == '9' {
												z01.PrintRune('\n')
											} else {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func PrintCombHuit() {
	for Mm := '0'; Mm <= '9'; Mm++ {
		for Mc := '0'; Mc <= '9'; Mc++ {
			for Md := '0'; Md <= '9'; Md++ {
				for Mu := '0'; Mu <= '9'; Mu++ {
					for m := '0'; m <= '9'; m++ {
						for c := '0'; c <= '9'; c++ {
							for d := '0'; d <= '9'; d++ {
								for u := '0'; u <= '9'; u++ {
									if Mm < Mc && Mc < Md && Md < Mu && Mu < m && m < c && c < d && d < u {
										z01.PrintRune(Mm)
										z01.PrintRune(Mc)
										z01.PrintRune(Md)
										z01.PrintRune(Mu)
										z01.PrintRune(m)
										z01.PrintRune(c)
										z01.PrintRune(d)
										z01.PrintRune(u)
										if Mm == '2' && Mc == '3' && Md == '4' && Mu == '5' && m == '6' && c == '7' && d == '8' && u == '9' {
											z01.PrintRune('\n')
										} else {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func PrintCombSept() {
	for Mc := '0'; Mc <= '9'; Mc++ {
		for Md := '0'; Md <= '9'; Md++ {
			for Mu := '0'; Mu <= '9'; Mu++ {
				for m := '0'; m <= '9'; m++ {
					for c := '0'; c <= '9'; c++ {
						for d := '0'; d <= '9'; d++ {
							for u := '0'; u <= '9'; u++ {
								if Mc < Md && Md < Mu && Mu < m && m < c && c < d && d < u {
									z01.PrintRune(Mc)
									z01.PrintRune(Md)
									z01.PrintRune(Mu)
									z01.PrintRune(m)
									z01.PrintRune(c)
									z01.PrintRune(d)
									z01.PrintRune(u)
									if Mc == '3' && Md == '4' && Mu == '5' && m == '6' && c == '7' && d == '8' && u == '9' {
										z01.PrintRune('\n')
									} else {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func PrintCombSix() {
	for Md := '0'; Md <= '9'; Md++ {
		for Mu := '0'; Mu <= '9'; Mu++ {
			for m := '0'; m <= '9'; m++ {
				for c := '0'; c <= '9'; c++ {
					for d := '0'; d <= '9'; d++ {
						for u := '0'; u <= '9'; u++ {
							if Md < Mu && Mu < m && m < c && c < d && d < u {
								z01.PrintRune(Md)
								z01.PrintRune(Mu)
								z01.PrintRune(m)
								z01.PrintRune(c)
								z01.PrintRune(d)
								z01.PrintRune(u)
								if Md == '4' && Mu == '5' && m == '6' && c == '7' && d == '8' && u == '9' {
									z01.PrintRune('\n')
								} else {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
	}
}

func PrintCombCinq() {
	for Mu := '0'; Mu <= '9'; Mu++ {
		for m := '0'; m <= '9'; m++ {
			for c := '0'; c <= '9'; c++ {
				for d := '0'; d <= '9'; d++ {
					for u := '0'; u <= '9'; u++ {
						if Mu < m && m < c && c < d && d < u {
							z01.PrintRune(Mu)
							z01.PrintRune(m)
							z01.PrintRune(c)
							z01.PrintRune(d)
							z01.PrintRune(u)
							if Mu == '5' && m == '6' && c == '7' && d == '8' && u == '9' {
								z01.PrintRune('\n')
							} else {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	}
}

func PrintCombQuatre() {
	for m := '0'; m <= '9'; m++ {
		for c := '0'; c <= '9'; c++ {
			for d := '0'; d <= '9'; d++ {
				for u := '0'; u <= '9'; u++ {
					if m < c && c < d && d < u {
						z01.PrintRune(m)
						z01.PrintRune(c)
						z01.PrintRune(d)
						z01.PrintRune(u)
						if m == '6' && c == '7' && d == '8' && u == '9' {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}

func PrintCombTrois() {
	for c := '0'; c <= '9'; c++ {
		for d := '0'; d <= '9'; d++ {
			for u := '0'; u <= '9'; u++ {
				if c < d && d < u {
					z01.PrintRune(c)
					z01.PrintRune(d)
					z01.PrintRune(u)
					if c == '7' && d == '8' && u == '9' {
						z01.PrintRune('\n')
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}

func PrintCombDeux() {
	for d := '0'; d <= '9'; d++ {
		for u := '0'; u <= '9'; u++ {
			if d < u {
				z01.PrintRune(d)
				z01.PrintRune(u)
				if d == '8' && u == '9' {
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}

func PrintCombUN() {
	for u := '0'; u <= '9'; u++ {
		z01.PrintRune(u)
		if u != '9' {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(10)
}
