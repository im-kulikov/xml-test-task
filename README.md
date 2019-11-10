# Тестовое задание по разбору XML

## Есть вот такой XML в кодировке windows-1251:

```xml
<?xml version="1.0" encoding="windows-1251"?>
<r title="корень">
    <a num="1"><st>одын</st><st>два</st></a>
    <a num="5"><st>три</st></a>
    <a num="27"><st>четыре</st></a>
</r>
```

Необходимо разобрать его потоково (важно!), каждый "a" положить в структуру (какую-нибудь, не важно)
и высести сам "a" с содержимым строкой.

## Под выводом имеется ввиду строки вида:
```
1: <a num="1"><st>одын</st><st>два</st></a>
....
2: <a num="27"><st>четыре</st></a>
```

## PS:

Собственно, приключения начинаются отсюда.
Зачем? Ну например вы хотите сравнивать и считаете контрольные суммы.

XML валидный.

## How to run

```
→ go run ./
0: <a num="1"><st>одын</st><st>два</st></a>
1: <a num="5"><st>три</st></a>
2: <a num="27"><st>четыре</st></a>
```
