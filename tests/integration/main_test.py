import unittest
import tempfile
import shutil
import json
# import pathlib
from os import path

from chaosad import chaos


# sys.path.insert(0, os.path.abspath('..'))
# from bootstrap.stuff import Stuff


class TestMain(unittest.TestCase):

    def setUp(self):
        self._tmpdir = tempfile.mkdtemp()

    def tearDown(self):
        shutil.rmtree(self._tmpdir)

    def test_should_create_output_file_properly(self):
        outputfn = path.join(self._tmpdir, "output-file")
        chaos.run(
            inputfn="/testdata/input-dump",
            outputfn=outputfn,
            listsize=3,
            concurrency=100,
        )

        count = 0
        # {"productId": "pid2", "images": ["http://www.xxx.com/3.png", "http://www.xxx.com/5.png", "http://www.xxx.com/6.png"]}
        with open(outputfn) as f:
            for line in f:
                j = json.loads(line)
                self.assertIn("productId", j)
                self.assertIn("images", j)
                count += 1
        self.assertGreater(count, 1)


if __name__ == '__main__':
    unittest.main()
