import random

letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
numbers = "0123456789"

def stringGenerate(size):
    return "".join(random.choice(letters + numbers) for i in range(size))
