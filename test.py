d1 = {'a': 100, 'b': 200, 'c': 300}
d2 = {'a': 300, 'b': 200, 'd': 400}
d = d1.copy()
for key, value in d2.items():
    d[key] = d.get(key, 0) + value
print(d)