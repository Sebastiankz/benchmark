import numpy as np
import time

def identityPrimo(num):
    if num < 2:
        return False
    for i in range(2, int(num**0.5) + 1):
        if num % i == 0:
            return False
    return True

def espiralMatrix(n):
    matrix = np.zeros((n, n), dtype = int)
    dx, dy = 0 , 1
    x, y = 0, 0
    for i in range(1, n*n + 1):
        matrix[x][y] = i
        if matrix[(x + dx) % n][(y + dy) % n]:
            dx, dy = dy, -dx
        x += dx
        y += dy
    return matrix

def addPrimeNumbers(matrix):
    sum = 0
    for row in matrix:
        for num in row:
            if identityPrimo(num):
                sum += num
    return sum

start = time.time()
N = 1000
createEspiralMatrix = espiralMatrix(N)

suma = addPrimeNumbers(createEspiralMatrix)

end = time.time()
execution_time = int((end - start) * 1000)
print(execution_time, "ms")
