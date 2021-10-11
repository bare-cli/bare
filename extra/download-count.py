from datetime import date


import requests as rq
import json


launchday = '2021-09-06'
today = date.today()
pkg_name = "barego"
URL = f"https://api.npmjs.org/downloads/point/{launchday}:{today}/{pkg_name}"

response = rq.get(URL)

if response.status_code == rq.codes.ok:
	download_count = str(json.loads(response.text)['downloads'])



f_content = f"""# Downloads\n[![SVG Banners](https://svg-banners.vercel.app/api?type=origin&text1=NPM&width=800&height=400)](https://www.npmjs.com/package/barego)
\n:calendar: Date: {today}, :calculator: Count: {download_count}
"""

with open("extra/download-count.md", "w") as f:
	f.write(f_content)
	