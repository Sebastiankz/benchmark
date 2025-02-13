#include <stdio.h>
#include <math.h>
#include <stdbool.h>

bool isPrime(int num)
{
    if (num < 2)
        return false;
    for (int i = 2; i <= sqrt(num); i++)
    {
        if (num % i == 0)
            return false;
    }
    return true;
}

void spiralMatrix(int n, int matrix[n][n])
{
    int dx = 0, dy = 1;
    int x = 0, y = 0;
    for (int i = 1; i <= n * n; i++)
    {
        matrix[x][y] = i;
        if (matrix[(x + dx + n) % n][(y + dy + n) % n] != 0)
        {
            int temp = dx;
            dx = dy;
            dy = -temp;
        }
        x += dx;
        y += dy;
    }
}

long long addPrimeNumbers(int n, int matrix[n][n])
{
    long long sum = 0;
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            if (isPrime(matrix[i][j]))
            {
                sum += matrix[i][j];
            }
        }
    }
    return sum;
}

int main()
{
    clock_t start = clock();

    int N = 1000;
    int matrix[N][N];
    spiralMatrix(N, matrix);

    long long suma = addPrimeNumbers(N, matrix);

    clock_t end = clock();
    double execution_time = (double)(end - start) / CLOCKS_PER_SEC * 1000;
    printf(execution_time + "ms\n");

    return 0;
}