orig_s = ""
with open("input.txt") as f:
    for line in f:
        orig_s += line.replace(" ", "").strip()

s = sorted(set(orig_s))
pairs = dict()

for ch in s:
    pairs[ch] = orig_s.count(ch)

m = max(pairs.values())

for row in range(m, 0, -1):
    s = ""
    for k in pairs:
        if pairs[k] % row == 0:
            s += "#"
            pairs[k] -= 1
        else:
            s += " "
    print(s)

print("".join(pairs.keys()))