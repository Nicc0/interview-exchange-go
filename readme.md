## Build

W projekcie znajduje się plik configuracyjny pod ścieżką `conf/app.ini.dist`. Należy umieścić w nim klucz do `openexchangeapi.org` i zmienić nazwę na `app.ini`. W razie konieczności można zmienić port, ale należy pamietać o przekazaniu go w dockerze.

> docker build --tag 'interview-exchange-go' .

> docker run --detach -p 8000:8000 'interview-exchange-go'

### Endpointy

```
GET /api/v1/rates
GET /api/v1/exchange
```

## Testy

Odnośnie testów... są one mocno prowizoryczne. Teoretycznie testami trzeba pokryć cały kod, a tutaj sprawdzamy tylko proste scenariusze. Musiałbym sprawdzić czy są jakieś fajne narzędzia które wizualnie pokazują pokrycie kodu i wtedy wyłapać szystkie edge case'y.

> go test

## Założenia i obserwacje

W ramach zadania rekrutacynego przyjąłem kilka założeń i uproszczeń:

- puste ciało odpowiedzi zrozumiałem jako pusty obiekt json
- wystawione endpointy posiadają prefix /api/v1 w celu lepszego uporządkowania
- kryptowaluty nie mogą być przewalutowane na zwykłe waluty
- przy pobieraniu kursów walut przydałoby się dodać cache, aby niepotrzebnie nie pytać w kółko o to samo i nie wykorzystywać zbędnie żądań do OpenExhnageAPI. Tutaj można by było napisać mechanizm który pobiera wszystkie dostępne kursy walut jeden raz (aktualizuje je gdy zajdzie taka potrzeba, tutaj potrzebna by była wiedza specjalistyczna, której na ten moment nie posiadam) a użytkownik dostaje odpowiedź z naszego serwera. Aktualny mechanizm działa w sumie jako proxy
- kursy kryptowalut są stałe i nie modyfikowalne podczas działania procesu API. Przy potencjalnie planowanym rozwoju należałoby skupić się na pobieraniu aktualnych kursów kryptowalut w czasie działania procesu np. pobieranie ich z zewnętrznego repozytorium (baza danych) np. przy każdym żądaniu lub gdy znamy interwały czasowe zmian kursu aktualizować dane w pamieci aplikacji
- z powodu duzej ilosci miejsc po przecinku, nalezalo uzyc niestandardowych typów (na szczescie jest biblioteka ktora dostarczona jest razem z Go, chodzi tutaj głównie o float który może operować na wiekszej ilości znaków po przeciunku)
- wyniki które są w przykładach nie są prawidłowe przez co trochę wpadłem w zakłopotanie i straciłem trochę czasu na weryfikacje np. wynik dla WBTC/USDT w pliku .md jest bliższy dla kursy 0.99 dla USDT/USD, dlatego cieżko mi tutaj jest zweryfikować na 100% poprawność wykonywania kodu
- przydałoby się również dodać zabezpieczenie przy zbyt dużych wartościach dla waluty, bilbioteka big ma również swoje limity, a przy pracy z walutami wiadomo jak jest