import scrapy

class BlibliSpider(scrapy.Spider):
    name = "blibli"
    allowed_domains = ["space.bilibili.com"]
    start_urls = [
        "https://space.bilibili.com/36802028/dynamic"
    ]
    def parse(self, response):

        pass
