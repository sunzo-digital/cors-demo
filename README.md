## Кейс 1

Простой запрос

Ответ будет заблокирован везде, кроме localhost
```javascript
fetch(
    'http://localhost:8080/without-any-headers', 
    {method: 'GET'}
).then(resp => resp.text()).then(console.log)
```

## Кейс 2

Простые запросы для разрешенного источника проходят
```javascript
fetch(
    'http://localhost:8080/only-simple',
    {method: 'GET'}
).then(resp => resp.text()).then(console.log) 
```

Сложные запросы не проходят, т.к. не обрабатываем preflight
```javascript
fetch(
    'http://localhost:8080/only-simple',
    {method: 'DELETE'}
).then(resp => resp.text()).then(console.log) 
```

## Кейс 3

Сложные запросы проходят для разрешенного источника
```javascript
fetch(
    'http://localhost:8080/',
    {method: 'DELETE'}
).then(resp => resp.text()).then(console.log) 
```

## Кейс 4

Разрешаем запросы с любых источников
```javascript
fetch(
    'http://localhost:8080/all-origins',
    {method: 'DELETE'}
).then(resp => resp.text()).then(console.log) 
```