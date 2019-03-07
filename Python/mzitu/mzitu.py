# encoding=utf-8

import requests
from bs4 import BeautifulSoup
from urllib.request import urlretrieve

download_links = []
path = ''
url = 'http://www.mzitu.com/'
headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 '
                  '(KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36'
}
res = requests.get(url, headers =headers)
soup = BeautifulSoup(res.text, 'lxml')
imgs = soup.select('li > a > img')

for img in imgs:
    print(img.get('data-original'))
    download_links.append(img.get('data-original'))

for item in download_links:
    urlretrieve(item, path+item[-5:])