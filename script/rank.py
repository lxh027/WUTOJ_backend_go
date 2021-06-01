import pandas as pd
import hashlib
import time
import datetime
import random
import string
import json
import requests

class Team:
    def __init__(self, teamName, member1, c1, member2, c2, member3, c3):
        self.teamName = ''
        self.teamMemberNames = []
        self.teamMemberClass = []
        self.teamMemberNames.append(member1)
        self.teamMemberClass.append(c1)
        self.teamMemberNames.append(member2)
        self.teamMemberClass.append(c2)
        self.teamMemberNames.append(member3)
        self.teamMemberClass.append(c3)
        self.teamName = teamName

    def __str__(self):
        return self.teamName + ' '.join(self.teamMemberNames) + ' '.join(self.teamMemberClass)

rankUrl = 'http://acmwhut.com/api/rank/contest/22'

problemsID = [46,47,48,49,50,51,52,53,54,55,56,57]

def getFinalRank():
    header = {'Content-Type': 'application/json'}
    requestInfo = requests.get(rankUrl, headers=header)
    # print(requestInfo.text)
    return json.loads(requestInfo.text)

rankInfo = getFinalRank()
# 队名 过题数 罚时 队员1 班级 队员2 班级 队员3 班级
# dataFrame
teamName = []
acmNumber = []
penalty = []
member1 = []
member2 = []
member3 = []
member1Class = []
member2Class = []
member3Class = []
# 队长姓名	队长性别	队长班级	队员姓名	队员性别	队员班级
teamInfo = pd.read_excel('./teaminfo.xlsx')
teamNames = teamInfo['队伍名字']
teamLeader = teamInfo['队长姓名']
teamLeaderClass = teamInfo['队长班级']
teamMembers = teamInfo['队员姓名']
teamMembersClass = teamInfo['队员班级']

teams = []
teamDic = {}

print(teamMembers)
print(teamLeader)
for i in range(0, int(len(teamNames)/2)):
    temp = Team(teamNames[2*i], teamLeader[2*i], teamLeaderClass[2*i], teamMembers[i*2], teamMembersClass[i*2], teamMembers[2*i+1], teamMembersClass[2*i+1])
    teams.append(temp)
    teamDic[teams[i].teamName] = teams[i]

# member1 = []
# member2 = []
# member3 = []
# member1Class = []
# member2Class = []
# member3Class = []

rankList = rankInfo['data']
for i in range(len(rankList)):
    if rankList[i]['nick'] == '志愿者':
        continue
    teamName.append(rankList[i]['nick'])
    acmNumber.append(rankList[i]['ac_num'])
    penalty.append(rankList[i]['penalty'])
    thisTeam = teamDic[rankList[i]['nick']]
    member1.append(thisTeam.teamMemberNames[0])
    member1Class.append(thisTeam.teamMemberClass[0])
    member2.append(thisTeam.teamMemberNames[1])
    member2Class.append(thisTeam.teamMemberClass[1])
    member3.append(thisTeam.teamMemberNames[2])
    member3Class.append(thisTeam.teamMemberClass[2])
    


outfile = pd.DataFrame({'队名': teamName, '过题数': acmNumber, '罚时': penalty, '队员1': member1, '队员1班级': member1Class, '队员2': member2, '队员2班级': member2Class,'队员3': member3, '队员3班级': member3Class})
outfile.to_csv('rank.csv')
