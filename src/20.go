import random

def gen_random_sequence(length):
    sequence = []
    while len(sequence) < length:
        index = random.randint(0, length - 1)
        if index not in sequence:
            sequence.append(index)
    return sequence

print(gen_random_sequence(5))
