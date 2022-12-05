/**
 * @param {string} s
 * @return {string}
 */
var frequencySort = function (s) {
	const m = new Map()
	for (char of s) m.set(char, m.get(char) + 1 || 1)

	const sortedMap = new Map([...m].sort((a, b) => b[1] - a[1]))
	let res = ""
	for (let item of sortedMap) {
		const char = item[0]
		const numOfChars = item[1]

		for (let i = 0; i < numOfChars; i++) {
			res += char
		}
	}

	return res
};