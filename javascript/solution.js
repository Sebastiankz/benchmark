function isPrime(num) {
    if (num < 2) return false;
    for (let i = 2; i <= Math.sqrt(num); i++) {
        if (num % i === 0) return false;
    }
    return true;
}

function spiralMatrix(n) {
    let matrix = Array.from({ length: n }, () => Array(n).fill(0));
    let dx = 0, dy = 1;
    let x = 0, y = 0;
    for (let i = 1; i <= n * n; i++) {
        matrix[x][y] = i;
        if (matrix[(x + dx + n) % n][(y + dy + n) % n]) {
            [dx, dy] = [dy, -dx];
        }
        x += dx;
        y += dy;
    }
    return matrix;
}

function addPrimeNumbers(matrix) {
    let sum = 0;
    for (let row of matrix) {
        for (let num of row) {
            if (isPrime(num)) {
                sum += num;
            }
        }
    }
    return sum;
}

const start = Date.now()

const N = 1000;

const createSpiralMatrix = spiralMatrix(N);
const suma = addPrimeNumbers(createSpiralMatrix);

const executionTime = Date.now() - start;
console.log(executionTime + 'ms');