import requests
import json
from urllib import request
from http import cookiejar
import pandas as pd

joinContestApi = 'http://acmwhut.com/api/contests/user/'

loginApi = 'http://acmwhut.com/api/login'
# 14?user_id=408&status=0
    
teamInfo = pd.read_csv('./out.csv')

contestID = '20'
status = 'status=0'
user_id = ''
teamName = teamInfo['账号']
teamPassword = teamInfo['密码']

headers = {'User-agent':'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36', 'Content-Type': 'application/json'}

def joinContest(nick, password):
    requestData = json.dumps({
        'password': password,
        'nick': nick,
        })
    response = requests.post(loginApi, headers=headers, data=requestData)
    responseDic  = json.loads(response.text)
    cookies = response.cookies
    joinUrl = joinContestApi + contestID + '?' + 'user_id=' + str(responseDic['data']['userId']) + '&' + status
    print(joinUrl)
    response = requests.post(joinUrl, headers=headers, cookies=cookies)
    print(response.text)

for i in range(0, len(teamName)):
    joinContest(teamName[i], teamPassword[i])