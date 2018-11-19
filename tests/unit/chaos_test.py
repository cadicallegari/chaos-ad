import unittest

from chaosad_py import main


# sys.path.insert(0, os.path.abspath('..'))
# from bootstrap.stuff import Stuff


class TestMain(unittest.TestCase):

    def test_should_sum_properly(self):
        self.assertEqual(4, main.sum(2, 2))


if __name__ == '__main__':
    unittest.main()
