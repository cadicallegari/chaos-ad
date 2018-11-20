import json
import asyncio
import logging

import aiohttp


logger = logging.getLogger(__name__)


async def fetch_url(client, url):
    async with client.get(url) as response:
        return response.status, await response.read()


def add_item(colection, key, item):
    items = colection.get(key, [])
    items.append(item)
    colection[key] = items


def dump_result(outputfn, products, urls):
    with open(outputfn, "w") as f:
        for k, v in products.items():
            f.write(json.dumps(
                {"productId": k, "images": v},
                ensure_ascii=False)
            )
            f.write("\n")


def should_process(products, urls, pid, url, maxsize):
    if url in urls or len(products.get(pid, [])) >= maxsize:
        return False
    return True


def item_handler_builder(fetcher, products, urls, maxsize):
    async def handler(client, item):
        p = json.loads(item)
        url = p.get("image")
        pid = p.get("productId")

        if not should_process(products, urls, pid, url, maxsize):
            return

        try:
            status, body = await fetcher(client, url)
            urls[url] = status == 200

        except aiohttp.ClientError as err:
            # TODO handle it properly
            print(err)
            return

        add_item(products, pid, url) if urls[url] else ""

    return handler


def consumer_builder(queue, handler):
    async def consumer(id):
        async with aiohttp.ClientSession() as client:
            while True:
                item = await queue.get()
                queue.task_done()
                if item is None:
                    break
                await handler(client, item)
    return consumer


def load_builder(inputfn):
    with open(inputfn) as f:
        for line in f:
            yield line


def producer_builder(queue, loader, concurrency):
    async def producer():
        for item in loader:
            await queue.put(item)

        for i in range(concurrency):
            await queue.put(None)

    return producer


async def _run(producer, consumer, concurrency):
    for i in range(concurrency):
        asyncio.ensure_future(consumer(i))

    await producer()
    # await queue.join()


def run(inputfn, outputfn, listsize, concurrency):
    loop = asyncio.get_event_loop()
    queue = asyncio.Queue(loop=loop, maxsize=concurrency)

    products = {}
    urls = {}

    loader = load_builder(inputfn)
    producer = producer_builder(queue, loader, concurrency)
    handler = item_handler_builder(fetch_url, products, urls, listsize)
    c_builder = consumer_builder(queue, handler)

    future = asyncio.ensure_future(
        _run(
            producer=producer,
            consumer=c_builder,
            concurrency=concurrency,
        )
    )

    loop.run_until_complete(future)

    dump_result(outputfn, products, urls)
