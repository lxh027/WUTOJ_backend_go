import pandas as pd
import hashlib
import time
import datetime
import random
import string
import json
import requests

omitSet = {'1', 'o', 'O', 'I', 'l', '0'}
charStr = string.ascii_lowercase + string.ascii_uppercase + '0123456789'

def getTeamName(teamNumber):
    number = str(teamNumber)
    teamName = 'team' + number.zfill(3) 
    return teamName

def getSeed():
    seed = ''
    for i in range(10):
        seed += charStr[random.randint(0, len(charStr)-1)]
    return seed

def getPassword(text):
    timestamp = int(time.time())
    res = str(datetime.datetime.fromtimestamp(timestamp))
    res = res + text
    hashCode = hashlib.md5()
    hashCode.update(res.encode("utf-8"))
    password = hashCode.hexdigest()
    password = password[:10]
    for i in range(0, len(password)):
        while password[i] in omitSet:
            password = password[:i] + charStr[random.randint(0, len(charStr)-1)] + password[i+1:]
    return password

registerApi = 'http://acmwhut.com/api/register'

def registerTeam(nick, realname, password):
    header = {'Content-Type': 'application/json'}
    requestData = json.dumps({
        'password': password,
        'password_check': password,
        'nick': nick,
        'realname': realname,
        'school': '武汉理工大学',
        'major': '比赛账号',
        'class': '比赛账号',
        'contact': '比赛账号',
        'mail':   'contest@whut.edu.cn' 
        })
    print(requestData)
    requestInfo = requests.post(registerApi, headers=header,data=requestData)
    print(requestInfo.text)
    

# teamInfo = pd.read_csv('./data.csv')
teamInfo = pd.read_excel('./teaminfo.xlsx')
teamID = []
teamPassword = []
teamNamesRaw = teamInfo['队伍名字']
teamNames = [teamNamesRaw[2*i] for i in range(0, int(len(teamNamesRaw)/2))]
print(len(teamNames))
print(teamNames)
random.shuffle(teamNames)
print(teamNames)
for i in range(0, len(teamNames)):
    teamID.append(getTeamName(i+1))
    teamPassword.append(getPassword(teamNames[i] + getSeed()))
    registerTeam(teamID[i], teamNames[i], teamPassword[i])

outfile = pd.DataFrame({"队名": teamNames, "账号": teamID, "密码": teamPassword})
outfile.to_csv("out.csv")


