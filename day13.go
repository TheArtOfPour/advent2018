package main

import (
	"fmt"
	"strings"
)

/*
Each time a cart has the option to turn (by arriving at any intersection),
it turns left the first time,
goes straight the second time,
turns right the third time,
and then repeats those directions starting again with left the fourth time
Move top left first
*/

type cart struct {
	nextTurn int // 0 left, 1 straight, 2 right
	x        int
	y        int
	track    rune
	crashed  bool
}

func getTracks(test string) [][]rune {
	stringSlice := strings.Split(test, "\n")
	tracks := make([][]rune, len(stringSlice))
	for i := 0; i < len(tracks); i++ {
		tracks[i] = make([]rune, len(stringSlice[0]))
	}
	for i, s := range stringSlice {
		s = strings.TrimRight(s, "\r")
		for j, char := range s {
			tracks[i][j] = char
		}
	}
	return tracks
}

func printTracks(tracks [][]rune) {
	fmt.Print("\n")
	for i := 0; i < len(tracks); i++ {
		for j := 0; j < len(tracks[i]); j++ {
			fmt.Printf("%s", string(tracks[i][j]))
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func getCarts(tracks [][]rune) []cart {
	carts := make([]cart, 0)
	for i := 0; i < len(tracks); i++ {
		for j := 0; j < len(tracks[i]); j++ {
			t := tracks[i][j]
			if t == '^' || t == 'v' {
				thisCart := cart{nextTurn: 0, x: i, y: j, track: '|', crashed: false}
				carts = append(carts, thisCart)
			} else if t == '<' || t == '>' {
				thisCart := cart{nextTurn: 0, x: i, y: j, track: '-', crashed: false}
				carts = append(carts, thisCart)
			}
		}
	}
	return carts
}

func moveCarts(tracks [][]rune, carts []cart) ([][]rune, []cart, bool) {
	newTracks := make([][]rune, len(tracks))
	for i := 0; i < len(newTracks); i++ {
		newTracks[i] = make([]rune, len(tracks[i]))
		for j := 0; j < len(tracks[i]); j++ {
			newTracks[i][j] = tracks[i][j]
		}
	}
	moved := make([]int, 0)
	// search tracks
	for j := 0; j < len(tracks[0]); j++ {
		for i := 0; i < len(tracks); i++ {
			t := tracks[i][j]
			// find cart
			if t == '^' || t == 'v' || t == '<' || t == '>' {
				for k := 0; k < len(carts); k++ {
					hasMoved := false
					for _, movedCart := range moved {
						if k == movedCart {
							hasMoved = true
							break
						}
					}
					if hasMoved {
						continue
					}
					if carts[k].x == i && carts[k].y == j {
						// nextTurn int // 0 left, 1 straight, 2 right
						nextTrack := ' '
						if t == '^' { // look up
							nextTrack = tracks[i-1][j]
							if nextTrack == '|' {
								newTracks[i-1][j] = '^'
							} else if nextTrack == '\\' {
								newTracks[i-1][j] = '<'
							} else if nextTrack == '/' {
								newTracks[i-1][j] = '>'
							} else if nextTrack == '+' {
								if carts[k].nextTurn == 0 {
									newTracks[i-1][j] = '<'
									carts[k].nextTurn++
								} else if carts[k].nextTurn == 1 {
									newTracks[i-1][j] = '^'
									carts[k].nextTurn++
								} else if carts[k].nextTurn == 2 {
									newTracks[i-1][j] = '>'
									carts[k].nextTurn = 0
								}
							} else if nextTrack == '>' || nextTrack == '<' || nextTrack == '^' || nextTrack == 'v' {
								fmt.Printf("CRASH! (%d,%d)\n", carts[k].y, carts[k].x)
								carts[k].crashed = true
								tracks[carts[k].x][carts[k].y] = carts[k].track
								newTracks[carts[k].x][carts[k].y] = carts[k].track
								for l := 0; l < len(carts); l++ {
									if l == k {
										continue
									}
									if carts[l].x == carts[k].x-1 && carts[l].y == carts[k].y {
										carts[l].crashed = true
										tracks[carts[l].x][carts[l].y] = carts[l].track
										newTracks[carts[l].x][carts[l].y] = carts[l].track
										break
									}
								}
							}
							carts[k].x--
						} else if t == 'v' { // look down
							nextTrack = tracks[i+1][j]
							if nextTrack == '|' {
								newTracks[i+1][j] = 'v'
							} else if nextTrack == '\\' {
								newTracks[i+1][j] = '>'
							} else if nextTrack == '/' {
								newTracks[i+1][j] = '<'
							} else if nextTrack == '+' {
								if carts[k].nextTurn == 0 {
									newTracks[i+1][j] = '>'
									carts[k].nextTurn++
								} else if carts[k].nextTurn == 1 {
									newTracks[i+1][j] = 'v'
									carts[k].nextTurn++
								} else if carts[k].nextTurn == 2 {
									newTracks[i+1][j] = '<'
									carts[k].nextTurn = 0
								}
							} else if nextTrack == '>' || nextTrack == '<' || nextTrack == '^' || nextTrack == 'v' {
								fmt.Printf("CRASH! (%d,%d)\n", carts[k].y, carts[k].x)
								carts[k].crashed = true
								tracks[carts[k].x][carts[k].y] = carts[k].track
								newTracks[carts[k].x][carts[k].y] = carts[k].track
								for l := 0; l < len(carts); l++ {
									if l == k {
										continue
									}
									if carts[l].x == carts[k].x+1 && carts[l].y == carts[k].y {
										carts[l].crashed = true
										tracks[carts[l].x][carts[l].y] = carts[l].track
										newTracks[carts[l].x][carts[l].y] = carts[l].track
										break
									}
								}
							}
							carts[k].x++
						} else if t == '<' { // look left
							nextTrack = tracks[i][j-1]
							if nextTrack == '-' {
								newTracks[i][j-1] = '<'
							} else if nextTrack == '\\' {
								newTracks[i][j-1] = '^'
							} else if nextTrack == '/' {
								newTracks[i][j-1] = 'v'
							} else if nextTrack == '+' {
								if carts[k].nextTurn == 0 {
									newTracks[i][j-1] = 'v'
									carts[k].nextTurn++
								} else if carts[k].nextTurn == 1 {
									newTracks[i][j-1] = '<'
									carts[k].nextTurn++
								} else if carts[k].nextTurn == 2 {
									newTracks[i][j-1] = '^'
									carts[k].nextTurn = 0
								}
							} else if nextTrack == '>' || nextTrack == '<' || nextTrack == '^' || nextTrack == 'v' {
								fmt.Printf("CRASH! (%d,%d)\n", carts[k].y, carts[k].x)
								carts[k].crashed = true
								tracks[carts[k].x][carts[k].y] = carts[k].track
								newTracks[carts[k].x][carts[k].y] = carts[k].track
								for l := 0; l < len(carts); l++ {
									if l == k {
										continue
									}
									if carts[l].x == carts[k].x && carts[l].y == carts[k].y-1 {
										carts[l].crashed = true
										tracks[carts[l].x][carts[l].y] = carts[l].track
										newTracks[carts[l].x][carts[l].y] = carts[l].track
										break
									}
								}
							}
							carts[k].y--
						} else if t == '>' { // look right
							nextTrack = tracks[i][j+1]
							if nextTrack == '-' {
								newTracks[i][j+1] = '>'
							} else if nextTrack == '\\' {
								newTracks[i][j+1] = 'v'
							} else if nextTrack == '/' {
								newTracks[i][j+1] = '^'
							} else if nextTrack == '+' {
								if carts[k].nextTurn == 0 {
									newTracks[i][j+1] = '^'
									carts[k].nextTurn++
								} else if carts[k].nextTurn == 1 {
									newTracks[i][j+1] = '>'
									carts[k].nextTurn++
								} else if carts[k].nextTurn == 2 {
									newTracks[i][j+1] = 'v'
									carts[k].nextTurn = 0
								}
							} else if nextTrack == '>' || nextTrack == '<' || nextTrack == '^' || nextTrack == 'v' {
								fmt.Printf("CRASH! (%d,%d)\n", carts[k].y, carts[k].x)
								carts[k].crashed = true
								tracks[carts[k].x][carts[k].y] = carts[k].track
								newTracks[carts[k].x][carts[k].y] = carts[k].track
								for l := 0; l < len(carts); l++ {
									if l == k {
										continue
									}
									if carts[l].x == carts[k].x && carts[l].y == carts[k].y+1 {
										carts[l].crashed = true
										tracks[carts[l].x][carts[l].y] = carts[l].track
										newTracks[carts[l].x][carts[l].y] = carts[l].track
										break
									}
								}
							}
							carts[k].y++
						}

						// update cart
						newTracks[i][j] = carts[k].track
						carts[k].track = nextTrack
						moved = append(moved, k)
						break
					}
				}

				// check for crashes
				for k := 0; k < len(carts)-1; k++ {
					if carts[k].crashed {
						continue
					}
					for l := k + 1; l < len(carts); l++ {
						if carts[l].crashed {
							continue
						}
						if carts[k].x == carts[l].x && carts[k].y == carts[l].y {
							fmt.Printf("CRASH! (%d,%d)\n", carts[k].y, carts[k].x)
							carts[k].crashed = true
							carts[l].crashed = true
							tracks[carts[l].x][carts[l].y] = carts[l].track
							newTracks[carts[l].x][carts[l].y] = carts[l].track
						}
					}
				}
			}
		}
	}
	newCarts := make([]cart, 0)
	for _, c := range carts {
		if !c.crashed {
			newCarts = append(newCarts, c)
		}
	}

	if len(newCarts) == 1 {
		fmt.Printf("ONE Cart Left! (%d,%d)\n", newCarts[0].y, newCarts[0].x)
	}

	return newTracks, newCarts, false
}

func advent13A(test string) int {
	tracks := getTracks(test)
	//printTracks(tracks)
	carts := getCarts(tracks)
	crashed := false
	for i := 0; i < 300; i++ {
		tracks, carts, crashed = moveCarts(tracks, carts)
		if crashed {
			fmt.Printf("\nCrashed!\n")
			break
		}
		//printTracks(tracks)
	}
	return 0
}

func advent13B(test string) int {
	tracks := getTracks(test)
	//printTracks(tracks)
	carts := getCarts(tracks)
	//crashed := false
	for i := 0; i < 100000; i++ {
		//printTracks(tracks)
		tracks, carts, _ = moveCarts(tracks, carts)
		if len(carts) == 1 {
			//tracks, carts, _ = moveCarts(tracks, carts)
			printTracks(tracks)
			fmt.Printf("\n%v\n", carts)
			break
		}
		//printTracks(tracks)
	}
	return 0
}
