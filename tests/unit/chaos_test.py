import unittest
# import tempfile
import asyncio
# import shutil

import aiounittest

from chaosad import chaos


# class TestMain(unittest.TestCase):

#     def setUp(self):
#         self._tmpdir = tempfile.mkdtemp()

#     def tearDown(self):
#         shutil.rmtree(self._tmpdir)

#     def test_should_sum_properly(self):
#         self.assertEqual(4, main.sum(2, 2))

class Test(unittest.TestCase):

    def setUp(self):
        self.loop = asyncio.new_event_loop()
        asyncio.set_event_loop(None)

    def test_xxx(self):
        async def go():
            ret = await chaos.add(5, 6)
            self.assertEqual(ret, 12)

        self.loop.run_until_complete(go())


class ChaosTest(aiounittest.AsyncTestCase):

    async def test_async_add(self):
        ret = await chaos.add(5, 6)
        self.assertEqual(ret, 11)

    # some regular test code
    def test_something(self):
        self.assertTrue(True)


if __name__ == '__main__':
    unittest.main()
