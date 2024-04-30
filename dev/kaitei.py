import sys
import math

count, radius = map(int, input().split()) ##火山の数と火山の半径を入力から読み取る
volcanoes=[]
for i in range(count):##count会の繰り返し処理で火山の座標を読み取る
    x,y= map(int, input().split())
    if 0 <= x < 100 and 0 <= y < 100:
        volcanoes.append((x, y))##リストに追加
    else:
        sys.exit(1)

def euclidean_distance(p1, p2):
    return math.sqrt((p1[0] - p2[0]) ** 2 + (p1[1] - p2[1]) ** 2) ##ユークリッド距離の計算

graph = {i: [] for i in range(len(volcanoes))}
for i in range(len(volcanoes)):
    for j in range(i + 1, len(volcanoes)):
        if euclidean_distance(volcanoes[i], volcanoes[j]) <= 2 * radius:
            graph[i].append(j)
            graph[j].append(i)


def dfs(node, visited, graph):
    visited[node] = True
    for neighbour in graph[node]:
        if not visited[neighbour]:
            dfs(neighbour, visited, graph)


visited = [False] * len(volcanoes)
island_count = 0
for i in range(len(volcanoes)):
    if not visited[i]:
        dfs(i, visited, graph)
        island_count += 1  

print(island_count)
