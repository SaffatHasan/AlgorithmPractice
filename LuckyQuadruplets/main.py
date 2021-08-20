import dp, bfs

def print_result(data):
    print(f"Running LuckyQuadruplets against {data=}")
    print(f"BFS got {bfs.LuckyQuadruplets(data)}")
    print(f"DP got {dp.LuckyQuadruplets(data)}")

data = [0,1,1,2,2,3,3,4,4,5,5,6,6,7,7,8,8,9,9]
current = data[:3]
for i in range(3, len(data)):
    current.append(data[i])
    print_result(current)

