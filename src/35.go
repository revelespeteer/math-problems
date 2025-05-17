def sum_of_squares(n):
    return n * (n + 1) * (2 * n + 1) // 6

import numpy as np

a = np.arange(10)
b = a**3
c = b.sum()
d = c.tolist()
print(d)
