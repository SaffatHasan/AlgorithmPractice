import math

def LuckyQuadruplets(tree: [int]) -> int:
    # assume tree is not null
    # assume tree is sorted
    total = 0
    children = getChildren(tree)
    # for i in range(len(tree)):
    #     print(f"children[{i}]={children[i]}")
    for node in range(len(tree)):
        for distance in range(int(math.log2(len(tree) + 1))):
            total += expand(tree, children, distance, node)
    return total

def expand(tree: [int], children: {int:[int]}, distance: int, node: int):
    parent = getParent(tree, node)
    # base case
    total = children[node][distance] if distance < len(children[node]) else 0

    # root does not look "up", only "down"
    if parent is None:
        return quadruplets(total)
        
    for offset in range(1, distance):
        # e.g. 2 away might look at parent's children excluding itself (node.parent[1] - current[0])
        #      or it  might look at grndpa's children excluding parent (node.parent.parent[0])
        if distance - offset == 0:
            total += 1
            break
        children_count = expandHelper(children, distance - offset, node, parent)
        if children_count is None:
            break

        # Don't update parent pointer if we're at root
        if getParent(tree, parent) is None:
            continue
        node = parent
        parent = getParent(tree, node)

    return quadruplets(total)


def expandHelper(children: {int:[int]}, distance: int, child: int, parent: int):
    parents_children = children[parent]
    childs_children = children[child]

    # cannot look down from parent
    # because distance is out of range
    if distance - 1 >= len(parents_children):
        return None

    choices = parents_children[distance-1]    

    # remove backtracking / duplicate nodes
    # i.e. node->parent->child is itself, is not 2 away
    if distance < len(childs_children):
        choices -= childs_children[distance]

    return choices

def quadruplets(number_of_choices):
    # fix current node as center
    # return all combinations of 3 for the children
    return math.comb(number_of_choices, 3)

def getChildren(tree: [int]) -> {int:[int]}:
    children = {}
    for i in range(len(tree)-1, 0, -1):
        if i in children:
            children[i][0] = 1
        else:
            children[i] = [1]

        parent = getParent(tree, i)
        if parent is None:
            continue

        if parent in children:
            children[parent] = mergeChildren(children[parent], children[i])
        else:
            children[parent] = mergeChildren([1], children[i])
    return children

def getParent(tree: [int], node: int) -> int:
    # root has no parent
    if tree[node] == 0:
        return None
    return tree[node]-1

def mergeChildren(parent: [int], child: [int]) -> [int]:
    """
    Given two lists of integers, merge the child into parent
    with 1 distance of offset (to account for parent being 1 away)
    """
    result = [0 for _ in range(max(len(parent), len(child)+1))]

    for i in range(len(parent)):
        result[i] = parent[i]

    for i in range(len(child)):
        result[i+1] += child[i]
    
    return result