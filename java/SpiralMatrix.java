
public class SpiralMatrix {

    public static boolean isPrime(int num) {
        if (num < 2) {
            return false;
        }
        for (int i = 2; i <= Math.sqrt(num); i++) {
            if (num % i == 0) {
                return false;
            }
        }
        return true;
    }

    public static int[][] spiralMatrix(int n) {
        int[][] matrix = new int[n][n];
        int dx = 0, dy = 1;
        int x = 0, y = 0;
        for (int i = 1; i <= n * n; i++) {
            matrix[x][y] = i;
            if (matrix[(x + dx + n) % n][(y + dy + n) % n] != 0) {
                int temp = dx;
                dx = dy;
                dy = -temp;
            }
            x += dx;
            y += dy;
        }
        return matrix;
    }

    public static long addPrimeNumbers(int[][] matrix) {
        long sum = 0;
        for (int[] row : matrix) {
            for (int num : row) {
                if (isPrime(num)) {
                    sum += num;
                }
            }
        }
        return sum;
    }

    public static void main(String[] args) {

        long startTime = System.currentTimeMillis();

        int N = 1000;
        int[][] createSpiralMatrix = spiralMatrix(N);
        long suma = addPrimeNumbers(createSpiralMatrix);

        long executionTime = System.currentTimeMillis() - startTime;
        System.out.println(executionTime + " ms");

    }
}
