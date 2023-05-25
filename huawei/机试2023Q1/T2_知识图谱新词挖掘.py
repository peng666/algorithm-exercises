# 给一串字母，找出里面的word有多少个
# 思路：从头到尾遍历字符串， 然后对word排序后比较就可以
content = input()
word = input()

cSize = len(content)
wSize = len(word)

if cSize < wSize:
    print(0)

newWord = "".join(sorted(list(word)))  # 变成有序的word

res = 0

i, j = 0, wSize
while j <= cSize:
    tmp = "".join(sorted(content[i:j]))
    if tmp == newWord:
        res += 1
    i += 1
    j += 1
print(res)
