const input1 = 5
const input2 = 3

function part1() {
	let p1 = input1
	let p2 = input2

	let s1 = 0
	let s2 = 0

	let i = 0
	for (; i < 1000; i++) {
		if (s1 >= 1000 || s2 >= 1000) {
			break
		}

		let add = (i + 1) * 3 * 3 - 3
		if (i % 2 == 0) {
			p1 += add
			p1 = p1 % 10
			s1 += p1 + 1
		} else {
			p2 += add
			p2 = p2 % 10
			s2 += p2 + 1
		}
	}

	console.log(i * 3 * Math.min(s1, s2))
	console.log("---")
}

part1()


function part2() {
	let win1 = 0
	let win2 = 0

	function play(i, dice, factor, p1, p2, s1, s2) {

		if (i % 2 == 0) {
			p1 += dice
			p1 = p1 % 10
			s1 += p1 + 1
		} else {
			p2 += dice
			p2 = p2 % 10
			s2 += p2 + 1
		}

		if (s1 >= 21) {
			win1 += factor
			return
		} else if (s2 >= 21) {
			win2 += factor
			return
		}

		play(i + 1, 3, factor * 1, p1, p2, s1, s2)
		play(i + 1, 4, factor * 3, p1, p2, s1, s2)
		play(i + 1, 5, factor * 6, p1, p2, s1, s2)
		play(i + 1, 6, factor * 7, p1, p2, s1, s2)
		play(i + 1, 7, factor * 6, p1, p2, s1, s2)
		play(i + 1, 8, factor * 3, p1, p2, s1, s2)
		play(i + 1, 9, factor * 1, p1, p2, s1, s2)
	}

	play(0, 3, 1, input1, input2, 0, 0)
	play(0, 4, 3, input1, input2, 0, 0)
	play(0, 5, 6, input1, input2, 0, 0)
	play(0, 6, 7, input1, input2, 0, 0)
	play(0, 7, 6, input1, input2, 0, 0)
	play(0, 8, 3, input1, input2, 0, 0)
	play(0, 9, 1, input1, input2, 0, 0)

	console.log(Math.max(win1, win2))
}

part2()
