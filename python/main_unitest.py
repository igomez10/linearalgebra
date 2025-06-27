import numpy as np
import unittest

class TestMatrixInverse(unittest.TestCase):
    def test_inverse(self):
        A = np.array([[1, 4], [2, 7]])
        Ainv = np.linalg.inv(A)
        identity = np.eye(2)
        result = A @ Ainv
        np.testing.assert_array_almost_equal(result, identity)
if __name__ == '__main__':
    unittest.main()
    print("All tests passed.")
    print("A * A inverse is identity matrix.")
    print("Test passed successfully.")
