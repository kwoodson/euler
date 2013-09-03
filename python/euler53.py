#!/usr/bin/env python


'''
 * There are exactly ten ways of selecting three from five, 12345:
 *
 * 123, 124, 125, 134, 135, 145, 234, 235, 245, and 345
 *
 * In combinatorics, we use the notation, 5C3 = 10.
 *
 * In general,
 *
 * nCr =        * n!
 * r!(nr)!
 * ,where r  n, n! = n(n1)...321, and 0! = 1.
 * It is not until n = 23, that a value exceeds one-million: 23C10 = 1144066.
 *
 * How many, not necessarily distinct, values of  nCr, for 1  n  100, are greater than one-million?
'''

def Fact(n):
    r = 1; 
    for i in range(2,n+1):
        r *= i
    return r


def main():

    count = 0
    for n in range(1,101):
        for r in range(1,n+1):
            #print "n,r=> %s,%s"%(n,r)
            fn = Fact(n)
            fr = Fact(r)
            fnr = Fact(n-r)
            z = fn / (fr * fnr)
            if(z > 1000000):
                count += 1;
    print "Count => %s" % count


if __name__ == "__main__":
    main()
