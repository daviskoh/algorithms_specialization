def zeroPad(numberString, zeros, left = True):
    """Return the string with zeros added to the left or right."""
    for i in range(zeros):
        if left:
            numberString = '0' + numberString
        else:
            numberString = numberString + '0'
    return numberString

def karatsubaMultiplication(x ,y):
    """Multiply two integers using Karatsuba's algorithm."""
    #convert to strings for easy access to digits
    x = str(x)
    y = str(y)
    #base case for recursion
    if len(x) == 1 and len(y) == 1:
        return int(x) * int(y)
    if len(x) < len(y):
        x = zeroPad(x, len(y) - len(x))
    elif len(y) < len(x):
        y = zeroPad(y, len(x) - len(y))

    n = len(x)
    j = n//2
    #for odd digit integers
    if (n % 2) != 0:
        print("ello !!!")
        print(n)
        print(j)
        j += 1

    a = int(x[:j])
    b = int(x[j:])
    c = int(y[:j])
    d = int(y[j:])

    #recursively calculate
    ac = karatsubaMultiplication(a, c)
    bd = karatsubaMultiplication(b, d)

    BZeroPadding = n - j
    AZeroPadding = BZeroPadding * 2

    k = karatsubaMultiplication(a + b, c + d)

    print("-------")
    # print((n - j) * 2)
    # print(n - j)
    # print(zeroPad(str(ac), AZeroPadding, False))
    print(n)
    print(j)
    print(x)
    print(y)
    print(a)
    print(b)
    print(c)
    print(d)
    # print(a+b)
    # print(c+d)
    # print(k)
    # print(str(k - ac - bd))
    # print(zeroPad(str(k - ac - bd), BZeroPadding, False))
    print("-------")

    A = int(zeroPad(str(ac), AZeroPadding, False))
    B = int(zeroPad(str(k - ac - bd), BZeroPadding, False))
    return A + B + bd

print(karatsubaMultiplication(123112, 1231))
