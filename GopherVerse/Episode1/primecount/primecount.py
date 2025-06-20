import math
import time

def is_prime(n):
    if n <= 1:
        return False
    sqrt_n = int(math.sqrt(n))
    for i in range(2, sqrt_n + 1):
        if n % i == 0:
            return False
    return True

start = time.perf_counter()
count = 0
for i in range(2, 10_000_001):
    if is_prime(i):
        count += 1
elapsed = (time.perf_counter() - start) * 1000
print(f"Total primes: {count}")
print(f"Execution time: {elapsed:.2f} ms")
