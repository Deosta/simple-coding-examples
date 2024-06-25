def fizzBuzz(self, n: int) -> List[str]:
    answer = []
    i = 1
    j = ""
    while i <= n:
        if i % 3 == 0 and i % 5 == 0:
            j = "FizzBuzz"
        elif i % 5 == 0:
            j = "Buzz"
        elif i % 3 == 0:
            j = "Fizz"
        else:
            j = str(i)
        answer.append(j)
        i += 1
    return answer