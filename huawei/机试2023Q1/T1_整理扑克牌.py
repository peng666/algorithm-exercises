# 根据给定的牌，按照牌型大小输出最大牌型
# 思路：用dict记录牌的数量、牌型的数量，一个个判断

cards = input().split()

# 整理扑克牌
# 牌牌记录器
cardsDict = {}
# 牌型记录器
cardStyleDict = {
    "danzhang": [],
    "duizi": [],
    "sanzhang": [],
    "zhadan": [],
    "hulu": []
}

# 遍历得到的牌--看看每张牌的数量
for card in cards:
    if cardsDict.get(card) is None:
        cardsDict[card] = 1
    else:
        cardsDict[card] += 1

# 遍历牌的数量，记录各种牌型
for card in cardsDict.keys():
    if cardsDict[card] == 1:
        cardStyleDict["danzhang"].append(card)
    elif cardsDict[card] == 2:
        cardStyleDict["duizi"].append(card)
    elif cardsDict[card] == 3:
        cardStyleDict["sanzhang"].append(card)
    else:
        cardStyleDict["zhadan"].append([card, cardsDict[card]])  # 数量也要存

# 排序
cardStyleDict["danzhang"].sort(reverse=True)
cardStyleDict["duizi"].sort(reverse=True)
cardStyleDict["sanzhang"].sort(reverse=True)
cardStyleDict["zhadan"].sort(reverse=True)

# 葫芦牌型--循环尝试组合
while len(cardStyleDict["sanzhang"]) > 0:
    # 对子用完了，并且3张只有1份，组合不出葫芦，退出
    if len(cardStyleDict["duizi"]) == 0 and len(cardStyleDict["sanzhang"]) == 1:
        break

    # 获取最大的三张
    sanzhang = cardStyleDict["sanzhang"].pop(0)
    tmp = None

    # 如果第二大的三张牌面，比第一大的对子牌面大；或者是没有对子；就用第二大的三张拆分来做葫芦
    if len(cardStyleDict["duizi"]) == 0 or (
            len(cardStyleDict["sanzhang"]) >= 1 and cardStyleDict["sanzhang"][0] > cardStyleDict["duizi"][0]):
        tmp = cardStyleDict["sanzhang"].pop(0)
        cardStyleDict["danzhang"].append(tmp)
    else:  # 否则就用对子来组合葫芦
        tmp = cardStyleDict["duizi"].pop(0)
    cardStyleDict["hulu"].append([sanzhang, tmp])

# 对单张牌再次排序，因为可能加了一些进去
cardStyleDict["danzhang"].sort(reverse=True)

# 结果
res = []

# 打印结果---按牌型大小输出
for zhadan, num in cardStyleDict["zhadan"]:
    res += [zhadan] * num
for sanzhang, duizi in cardStyleDict["hulu"]:
    res += [sanzhang] * 3 + [duizi] * 2
for sanzhang in cardStyleDict["sanzhang"]:
    res += [sanzhang] * 3
for duizi in cardStyleDict["duizi"]:
    res += [duizi] * 2
for danzhang in cardStyleDict["danzhang"]:
    res += [danzhang]

print(" ".join(map(str, res)))
