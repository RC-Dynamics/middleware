from threading import Lock, Thread
from time import sleep
from random import randint

lock_A = Lock()
lock_B = Lock()
a = 10
b = 3

def fsum():
    print "init sum thread"
    while (not lock_A.acquire()):
        print "fsum fail lock A"
        sleep(randint(1,5))
    sleep(2)
    print "fsum lock A"
    while (not lock_B.acquire()):
        print "fsum fail lock B"
        sleep(randint(1,5))
    print "fsub lock B"

    print(a+b)
    lock_B.release()
    lock_A.release()
    print "finish sum thread"

def fsub():
    print "init sub thread"
    while (not lock_A.acquire()):
        print "fsub fail lock A"
        sleep(randint(1,5))
    print "fsum lock A"
    while (not lock_B.acquire()):
        print "fsub fail lock B"
        sleep(randint(1,5))
    sleep(2)
    print "fsub lock B"

    print(a-b)
    lock_A.release()
    lock_B.release()
    print "finish sub thread"

if __name__ == "__main__":
    threads = []
    for func in [fsum, fsub]:
        threads.append(Thread(target=func))
        threads[-1].start()

    for thread in threads:
        thread.join()
