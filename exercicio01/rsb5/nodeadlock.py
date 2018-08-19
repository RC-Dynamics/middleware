from threading import Lock, Thread
from time import sleep
from random import randint

lock_A = Lock()
lock_B = Lock()
a = 10
b = 3

def fsum():
    tries = 0
    print "init sum thread"
    while (not lock_A.acquire(False)):
        print "fsum fail lock A"
        sleep(randint(1,5))
    sleep(2)
    print "fsum lock A"
    while (not lock_B.acquire(False)):
        print "fsum fail lock B"
        tries+=1
        if(tries>=5):
            lock_A.release()
            print "fsum unlock A"
            sleep(6)
            fsum()
            return
        sleep(randint(1,5))

    print(a+b)
    lock_B.release()
    lock_A.release()
    print "finish sum thread"

def fsub():
    tries = 0
    print "init sub thread"
    while (not lock_B.acquire(False)):
        print "fsub fail lock B"
        sleep(randint(1,5))
    sleep(2)
    print "fsub lock B"
    while (not lock_A.acquire(False)):
        print "fsub fail lock A"
        tries+=1
        if(tries>=5):
            lock_B.release()
            print "fsub unlock B"
            sleep(6)
            fsub()
            return
        sleep(randint(1,5))

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
