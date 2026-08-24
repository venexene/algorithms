# Balanced Binary Tree

## Решение

### DFS снизу вверх — O(n), O(h)

`dfs(node)` возвращает пару `(сбалансировано, высота)`:

- `nil` → `(true, 0)`;
- рекурсивно считаем `(leftFlag, leftDepth)` и `(rightFlag, rightDepth)`;
- `flag = |leftDepth - rightDepth| <= 1`;
- возвращаем `(leftFlag && rightFlag && flag, max(leftDepth, rightDepth) + 1)`.

Ответ — первое значение для корня.

## Что важно запомнить

- **Два возвращаемых значения** — флаг и высота: пост-обход, сначала дети, потом узел.
- **Сбалансированность узла** = `|leftDepth - rightDepth| <= 1` **и** оба поддерева сбалансированы.
- **Высота** = `max(leftDepth, rightDepth) + 1` — та же механика, что в Diameter of Binary Tree.
- **Пустое дерево** сбалансировано: `nil` → `(true, 0)`.
- Модуль разности можно считать через `max(a-b, b-a)` или `abs(left-right)`.