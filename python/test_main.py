import unittest

import numpy as np
import pandas as pd


class TestMatrixInverse(unittest.TestCase):
    def test_inverse(self):
        A = np.array([[1, 4], [2, 7]])
        Ainv = np.linalg.inv(A)
        identity = np.eye(2)
        result = A @ Ainv
        np.testing.assert_array_almost_equal(result, identity)


class TestPCACustom(unittest.TestCase):
    file = "./data/pca_dataset.csv"
    # read data and do pca
    df = pd.read_csv(file)
    
    # do pca
    from sklearn.decomposition import PCA
    pca = PCA(n_components=2)
    data = df.values
    pca.fit(data)
    
    # print the principal components
    print("Principal Components:")
    for i in range(2):
        print(f"Component {i+1}:")
        print(f"Vector: {pca.components_[i]}")
        print(f"Variance: {pca.explained_variance_[i]}")
    
    def test_pca(self):
        # check that the explained variance is non-negative
        self.assertTrue(all(self.pca.explained_variance_ >= 0))
        
        # check that the components are orthogonal
        dot_product = np.dot(self.pca.components_[0], self.pca.components_[1])
        self.assertAlmostEqual(dot_product, 0, places=5)

if __name__ == "__main__":
    unittest.main()
    print("All tests passed.")
    print("A * A inverse is identity matrix.")
    print("Test passed successfully.")
