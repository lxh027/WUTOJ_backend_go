import urllib.request
import json
import sys
# import getpot
wqy = sys.argv
str = ""
print("wqy")
# with open('../commit.log', 'r') as f:
# 	wcr = f.readlines()
# flag = 1
# for i in wcr:
# 	if i[0:6] == "Author":
# 		ljw = i.split("<")
# 		i = ljw[0]
# 		flag = 0
# 	if flag:
# 		continue
# 	if len(i.strip()) < 3:
# 		continue
# 	str += i.strip()+"\n"
# # f.close()
print("wyh")
data = {
	"ToUserUid": int(wqy[1]),
	"SendToType": 2,
	"SendMsgType": "TextMsg",
	"Content": "yzb又瞎push了什么怪东西进https://github.com/UncleMadeleine/WUTOJ_backend_go/commits/main：\n"+str
}
print("ljw")
values = urllib.parse.urlencode(data).encode(encoding='UTF8')
headers = {'Content-Type': 'application/json'}
print(data)
print(values)
print(json.dumps(data))
print(json.dumps(data).encode())
request = urllib.request.Request(
	url=wqy[2], headers=headers, data=json.dumps(data).encode())
response = urllib.request.urlopen(request)
