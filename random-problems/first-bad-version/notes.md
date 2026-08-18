# First Bad Version

## Решение

### Бинарный поиск «первого true» — O(log n), O(1)

Ищем первую плохую версию в монотонной последовательности `false, false, ..., true, true`. Сужаем `[left, right]`: если `isBadVersion(mid) == false` — `left = mid + 1`, иначе `right = mid`. В конце `left == right` — это и есть ответ.

## Что важно запомнить

- **Паттерн «найти первое true»** — классика бинарного поиска по предикату. Шаблон: `for left < right`, `mid := left + (right-left)/2`, при `false` двигать `left`, при `true` — `right = mid` (не `mid-1`, чтобы не потерять первое `true`).
- **`left := 1`, а не 0** — версии нумеруются с 1; вызов `isBadVersion(0)` вне контракта API.
- **Overflow-safe `mid`** — `left + (right-left)/2` вместо `(left+right)/2`, чтобы не переполнить при больших `left+right`.
- **`return right`** — в конце цикла `left == right`, возвращать можно любой из них.
- Для теста стаб задаётся порогом: `var bad = 4; isBadVersion(v) = v >= bad`.

## Что пошло не так

1. **`mid := (left + right) / 2`** — риск переполнения; заменено на `left + (right-left)/2`.