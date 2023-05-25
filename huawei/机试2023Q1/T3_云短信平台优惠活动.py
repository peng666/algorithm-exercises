# 给出客户预算、短信价目表充1元获得多少短信。求用户获得最多的短信数
# 动态规划。有点像背包问题

money = int(input())
prices = list(map(int, input().split()))

dp = [0 for i in range(money + 1)]  # 多少钱最多的短信数
for i in range(len(prices) + 1):  # 遍历所有的短信套餐
    for j in range(money + 1):  # 遍历所有的预算
        if i == 0:  # 初始状态：预算为0，短信数为0
            continue
        if j == 0:  # 初始状态，套餐的短信数为0，那就是0
            continue
        if j > i:  # 初始状态：不够钱，那就是0
            continue
        if dp[j] < dp[j - i] + prices[i - 1]:  # 贪心思想
            dp[j] = dp[j - 1] + prices[i - 1]  # 这种情况的短信数更多，那就这个
    print(dp[money])
