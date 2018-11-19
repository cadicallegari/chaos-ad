import sys
import logging
import importlib

from scrapy.utils.project import get_project_settings

from blackpearl import settings
from blackpearl import loop
from tortuga import argsparser
from tortuga.log import config as logconfig


def setup_crawler(
        spider_module,
        spider_class_name,
        proxy_enabled,
        stats_collector,
        concurrency
):
    spider_module = importlib.import_module(spider_module)
    SpiderClass = getattr(spider_module, spider_class_name)
    settings = get_project_settings()
    settings.set('PROXY_ENABLED', proxy_enabled)
    settings.set('STATS_COLLECTOR_TYPE', stats_collector)
    loop.run(SpiderClass, settings, concurrency)


if __name__ == '__main__':

    parser = argsparser.new(description='Starts Black Pearl Daemon')
    parser.add_argument('--crawler', help='Crawler to be executed')
    args = parser.parse_args()

    if args.crawler is None:
        print("--crawler is obligatory. run --help for more details")
        sys.exit()

    print("blackpearld.init: configuring log level: {}".format(args.log_level))
    logconfig.config(args.log_level)
    log = logging.getLogger("blackpearld.init")
    log.info("configured log level to '{}' with success. Starting loop".format(
        args.log_level
    ))

    spider_class = args.crawler.split(".")[-1]
    spider_module = args.crawler.split(".")[:-1]
    spider_module = ".".join(spider_module)

    settings.STATS_COLLECTOR_TYPE = args.stats_collector
    settings.STORAGE_PIPELINE_ENABLED = args.enable_storage_pipeline
    settings.DEBUG_STORAGE_ENABLED = args.enable_debug_storage
    settings.DEBUG_STORAGE_DIR = args.debug_storage_dir

    if args.enable_proxy:
        if not args.proxy_loaders:
            print("--proxy-loaders is required with proxy enabled."
                "If you don't want to use proxy set --no-proxy "
                "run --help for more details")
            sys.exit()
        log.info("raw proxy loaders: {}".format(args.proxy_loaders))
        settings.PROXY_LOADERS = args.proxy_loaders.split(",")
        log.info("parsed proxy loaders: {}".format(settings.PROXY_LOADERS))

    if args.download_timeout is not None:
        settings.DOWNLOAD_TIMEOUT = int(args.download_timeout)
        log.info("download timeout: {}".format(settings.DOWNLOAD_TIMEOUT))

    setup_crawler(
        spider_module,
        spider_class,
        args.enable_proxy,
        args.stats_collector,
        args.concurrency
    )
