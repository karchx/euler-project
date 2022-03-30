limit = 4_000_000
aux = 1
fib = 0
result = 0

while aux <= limit:
    aux += fib
    fib = aux - fib
    if fib%2 == 0:
        result += fib
print(result)
