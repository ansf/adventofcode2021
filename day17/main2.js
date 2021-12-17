// let's just brute force it

const targetX = [56, 76]
const targetY = [-162, -134]

const maxDy = -targetY[0] -1
const maxDx = targetX[1]


count = 0
for (let dy = maxDy; dy >= targetY[0]; dy--) {
	for (let dx = maxDx; dx >= 0; dx--) {
		if (isHit(dx, dy)) {
			count++
		}
	}
}
console.log(count)

function isHit(dx, dy) {
	const initialDx = dx
	const initialDy = dy

	let y = 0	
	let x = 0
	while (true) {
		y += dy
		x += dx

		dy--
		if (dx > 0) {
			dx--
		}

		if (y >= targetY[0] && y <= targetY[1]
			&& x >= targetX[0] && x <= targetX[1]) {
			return true
		} else if (y < targetY[0] || x > targetX[1]) {
			return false
		}
	}
}
