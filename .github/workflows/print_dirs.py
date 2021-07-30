import os
import json


def get_dirs():
    exclude = ['.git', '.github']
    return [x for x in os.listdir() if x not in exclude and os.path.isdir(x)]

def to_json(data):
    return json.dumps(data)

if __name__ == "__main__":
    print(to_json(get_dirs()))