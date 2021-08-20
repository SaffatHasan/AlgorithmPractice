import math

def LuckyQuadruplets(tree: [int]):
    graph = makeGraph(tree)

    # print(graph)

    result = 0
    for i in range(len(tree)):
        result += bfs(i, graph)
    return result

def bfs(node: int, graph: {int: [int]}):
    visited = set()
    visited.add(node)

    # print(f"\nExpanding at node {node}")

    # initially populate queue with nodes that are 1 distance away
    # i.e. the neighbors of the current node
    queue = [neighbor for neighbor in graph[node]]
    result = 0
    distance = 0
    while len(visited) < len(graph) and len(queue) > 0:
        distance += 1
        # print(f"{queue=} for {distance=}")
        # each layer in BFS represents
        # the number of nodes that are equidistant to
        # the current node
        layer_size = len(queue)
        if layer_size >= 3:
            # print(f"{layer_size=} adding {quadruplets(layer_size)=} quadruplets...")
            result += quadruplets(layer_size) # number of quadruplets we can make with current layer

        # expand
        for _ in range(layer_size):
            current_node = queue[0]
            queue = queue[1:]
            visited.add(current_node)

            for neighbor in graph[current_node]:
                if neighbor in visited:
                    continue
                queue.append(neighbor)
        
    return result

def quadruplets(node_count):
    return math.comb(node_count, 3)

def makeGraph(tree: [int]):
    result = {i: [] for i in range(len(tree))}

    for node in range(len(tree)):
        parent = getParent(tree, node)

        if parent is None:
            continue
        result[node].append(parent)
        result[parent].append(node)

    return result

def getParent(tree: [int], node: int) -> int:
    if tree[node] == 0:
        # root has no parent
        return None

    return tree[node] - 1
