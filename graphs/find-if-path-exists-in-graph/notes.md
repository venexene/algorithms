# Find if Path Exists in Graph

## Решение

### 1. DFS + adjacency list + visited — O(V + E), O(V + E)
Строим adjacency list: `adj[vertex] = []neighbors`. DFS от source, visited метки на узлах. Если дошли до destination — true.

### 2. Union Find (Disjoint Set Union) — O(V + E), O(V)
Каждая вершина — своё множество. `find(x)` находит корень, `union(x, y)` объединяет множества. В конце: `find(source) == find(destination)`.

## Что пошло не так (7 итераций)

### 1. Бесконечный цикл (TLE)
DFS без visited. Граф с циклами → рекурсия бесконечно ходит по кругу. Нужен `visited map[int]bool`.

### 2. Visited на рёбрах, не на узлах
`visited[[2]int{edge[0], edge[1]}]` — не работает. Узел может быть посещён через разные рёбра, отметка ребра не предотвращает повторный вход в узел. Нужен `visited[node]`.

### 3. O(V × E) — TLE на больших графах
DFS сканировал ВСЕ рёбра для каждого узла. При 10⁵ узлов и 2·10⁵ рёбер — квадратичная сложность. Нужен adjacency list: построить один раз за O(E), дальше соседи за O(1).

### 4. Adjacency list построен только в одну сторону
`adj[edge[0]] = append(adj[edge[0]], edge[1])` без обратного направления. Граф двунаправленный — нужно добавлять оба.

### 5. Union без find
`parent[edge[1]] = parent[edge[0]]` — не учитывает цепочки. Если `parent[edge[0]]` сам ещё не корень, связь теряется. Нужна функция `find`, идущая до корня.

### 6. Бесконечный цикл в find
`for parent[node] != node { parent[node] = find(...) }` — цикл не выходит после сжатия. Нужно `if`, не `for`.

### 7. TLE без path compression
`find` без сжатия пути — цепочка из 100k вершин даёт O(n) на каждый вызов. С path compression: `parent[node] = find(parent[node])` делаем всех на пути указывать на корень.

## Что важно запомнить

### Adjacency list
- Строится один раз за O(E) из списка рёбер
- Для неориентированного графа — добавлять оба направления
- `adj := make([][]int, n)` если вершины 0..n-1, иначе мапа

### DFS на графе
- visited на УЗЛАХ, не на рёбрах
- Проверка visited ДО рекурсии
- `res = res || dfs(neighbor)` — short-circuit, выходит рано

### Union Find
- `find(x)` — рекурсивно идёт к корню
- `union(x, y)` — `root1, root2 := find(x), find(y); parent[root2] = root1`
- Path compression: `parent[node] = find(parent[node])` — делает дерево плоским
- Инициализация: `for i := 0; i < n; i++ { parent[i] = i }` — каждая вершина свой корень
- Проверка: `find(source) == find(destination)`