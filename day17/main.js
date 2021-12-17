const target = [-162, -134]

// we want to be at y=0 with a velocity lower than -162 to still hit the target zone

let dy = 161
let y = 0
let maxY = 0

while (true) {
	y += dy
	dy -= 1

	maxY = Math.max(y, maxY)


	if (y < -162) {
		break
	}
}

console.log(maxY)
