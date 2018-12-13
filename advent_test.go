package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDay1(t *testing.T) {
	input := `+1
-2
+3
+1`
	out, _ := advent1A(input)
	expected := 3
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `+1
-1`
	out, _ = advent1B(input)
	expected = 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `+3
+3
+4
-2
-4`
	out, _ = advent1B(input)
	expected = 10
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `-6
+3
+8
+5
-6`
	out, _ = advent1B(input)
	expected = 5
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `+7
+7
-2
-7
-4`
	out, _ = advent1B(input)
	expected = 14
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay2(t *testing.T) {
	input := `abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`
	out, _ := advent2A(input)
	expected := 12
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`
	out2, _ := advent2B(input)
	expected2 := `fgij`
	if !cmp.Equal(out2, expected2) {
		t.Errorf("Didn't match %s", cmp.Diff(out2, expected2))
	}
}

func TestDay3(t *testing.T) {
	input := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`
	out, _ := advent3A(input)
	expected := 4
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	out, _ = advent3B(input)
	expected = 3
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay4(t *testing.T) {
	input := `[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`
	out, _ := advent4A(input)
	expected := 240
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	out, _ = advent4B(input)
	expected = 4455
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay5(t *testing.T) {
	input := `dabAaCBAcCcaDA`
	out, _ := advent5A(input)
	expected := 10
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	out, _ = advent5B(input)
	expected = 4
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay6(t *testing.T) {
	input := `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`
	out, _ := advent6A(input)
	expected := 17
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	out, _ = advent6B(input)
	expected = 16
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay7(t *testing.T) {
	input := `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`
	// out, _ := advent7A(input)
	// expected := "CABDFE"
	// if !cmp.Equal(out, expected) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	// }
	outInt, _ := advent7B(input)
	expectedInt := 15
	if !cmp.Equal(outInt, expectedInt) {
		t.Errorf("Didn't match %s", cmp.Diff(outInt, expectedInt))
	}
	// 1175 TOO LOW
	// !1235
	// !1251
}

func TestDay8(t *testing.T) {
	input := `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`
	out := advent8A(input)
	expected := 138
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	out = advent8B(input)
	expected = 66
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay9(t *testing.T) {
	input := `9 players; last marble is worth 25 points`
	out := advent9A(input)
	expected := 32
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `10 players; last marble is worth 1618 points`
	out = advent9A(input)
	expected = 8317
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `13 players; last marble is worth 7999 points`
	out = advent9A(input)
	expected = 146373
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `17 players; last marble is worth 1104 points`
	out = advent9A(input)
	expected = 2764
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `21 players; last marble is worth 6111 points`
	out = advent9A(input)
	expected = 54718
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `30 players; last marble is worth 5807 points`
	out = advent9A(input)
	expected = 37305
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	// out = advent9B(input)
	// expected = 0
	// if !cmp.Equal(out, expected) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	// }
}

func TestDay10(t *testing.T) {
	input := `position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>`
	out := advent10A(input)
	expected := 1
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	// out = advent10B(input)
	// expected = 0
	// if !cmp.Equal(out, expected) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	// }
}

func TestDay11(t *testing.T) {
	unitTest := calculatePower(3, 5, 8)
	unitExpected := 4
	if !cmp.Equal(unitTest, unitExpected) {
		t.Errorf("Didn't match %s", cmp.Diff(unitTest, unitExpected))
	}
	unitTest = calculatePower(122, 79, 57)
	unitExpected = -5
	if !cmp.Equal(unitTest, unitExpected) {
		t.Errorf("Didn't match %s", cmp.Diff(unitTest, unitExpected))
	}
	unitTest = calculatePower(217, 196, 39)
	unitExpected = 0
	if !cmp.Equal(unitTest, unitExpected) {
		t.Errorf("Didn't match %s", cmp.Diff(unitTest, unitExpected))
	}
	unitTest = calculatePower(101, 153, 71)
	unitExpected = 4
	if !cmp.Equal(unitTest, unitExpected) {
		t.Errorf("Didn't match %s", cmp.Diff(unitTest, unitExpected))
	}
	// input := `18`
	// out := advent11A(input)
	// expected := 138
	// if !cmp.Equal(out, expected) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	// }
	// out = advent11B(input)
	// expected = 66
	// if !cmp.Equal(out, expected) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	// }
}

func TestDay12(t *testing.T) {
	input := `initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`
	out := advent12A(input)
	expected := 325
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	// out = advent12B(input)
	// expected = 66
	// if !cmp.Equal(out, expected) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	// }
}

func TestAdvent(t *testing.T) {
}
