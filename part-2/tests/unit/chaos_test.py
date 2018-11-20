import unittest
import asyncio


from chaosad import chaos


class ChaosTest(unittest.TestCase):

    def setUp(self):
        self.loop = asyncio.new_event_loop()
        asyncio.set_event_loop(None)

    def test_should_add_url_to_product_on_success(self):
        async def fakefetcher(c, url):
            return 200, ""

        prods = {}
        urls = {}
        handler = chaos.item_handler_builder(fakefetcher, prods, urls, 1)

        raw = '{"productId": "pid", "image":"fake/url"}'

        self.loop.run_until_complete(handler(None, raw))

        self.assertEqual({"pid": ["fake/url"]}, prods)
        self.assertEqual({"fake/url": True}, urls)

    def test_should_respect_url_limit(self):
        async def fakefetcher(c, url):
            return 200, ""

        prods = {}
        urls = {}
        handler = chaos.item_handler_builder(fakefetcher, prods, urls, 2)

        for i in range(5):
            raw = '{"productId": "pid", "image":"fake/url-%d"}' % i
            self.loop.run_until_complete(handler(None, raw))

        self.assertEqual(
            {'pid': ['fake/url-0', 'fake/url-1']},
            prods
        )

    def test_should_not_add_url_that_not_return_ok(self):
        async def fakefetcher(c, url):
            if "fail" in url:
                return 404, ""
            return 200, ""

        prods = {}
        urls = {}
        handler = chaos.item_handler_builder(fakefetcher, prods, urls, 2)

        for i, s in enumerate(["fail", "success", "success", "fail", "success"]):
            raw = '{"productId": "pid", "image":"fake/url-%i-%s"}' % (i, s)
            self.loop.run_until_complete(handler(None, raw))

        self.assertEqual(
            {'pid': ['fake/url-1-success', 'fake/url-2-success']},
            prods
        )


if __name__ == '__main__':
    unittest.main()
