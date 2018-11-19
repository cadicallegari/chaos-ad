import json
import asyncio
import logging

import aiohttp


logger = logging.getLogger(__name__)


async def add(x, y):
    await asyncio.sleep(0.1)
    return x + y


async def hit_url(client, url):
    async with client.get(url) as response:
        return response.status, await response.read()


async def news_producer(queue, inputfn, concurrency):
    with open(inputfn) as f:
        for line in f:
            await queue.put(line)

    for i in range(concurrency):
        await queue.put(None)


def add_item(colection, key, item):
    items = colection.get(key, [])
    items.append(item)
    colection[key] = items


async def news_consumer(id, queue, products, urls, listsize):
    async with aiohttp.ClientSession() as client:
        while True:
            item = await queue.get()
            queue.task_done()
            if item is None:
                break

            # {"productId":"pid482","image":"http://localhost:4567/images/167410.png"}
            p = json.loads(item)
            url = p.get("image")

            try:
                if url not in urls:
                    status, body = await hit_url(client, url)
                    urls[url] = status == 200
            except aiohttp.ClientError as err:
                # TODO handle it properly
                print(err)
                continue

            r = urls[url]
            if r:
                add_item(products, p.get("productId"), url)


async def _run(queue, inputfn, products, urls, listsize, concurrency):
    for i in range(concurrency):
        asyncio.ensure_future(
            news_consumer(i, queue, products, urls, listsize)
        )

    await news_producer(queue, inputfn, concurrency)
    await queue.join()


def _dump_result(outputfn, products, urls):
    print(outputfn)
    print(urls)
    print(products)


def run(inputfn, outputfn, listsize, concurrency):
    loop = asyncio.get_event_loop()
    queue = asyncio.Queue(loop=loop, maxsize=concurrency)

    products = {}
    urls = {}

    future = asyncio.ensure_future(
        _run(
            queue=queue,
            inputfn=inputfn,
            products=products,
            urls=urls,
            listsize=listsize,
            concurrency=concurrency,
        )
    )

    loop.run_until_complete(future)

    _dump_result(outputfn, products, urls)
