# Napredne-tehnike-programiranja
Web aplikacija za pronalazak majstora kucnih poslova. 

## Funkcionalnosti

1. Neregistrovani korisnik
- Može da se registruje, putem mail-a dobija aktivacioni link
- Pregled, pretraga i sortiranje majstora po zanimanju i slobodnim terminima (radnom vremenu), cijeni usluge, ocjenama, mjestu prebivalista, imenu 

2. Registrovani korisnik
- Pregled profila i izmjena osnovnih podataka
- Rezervacija termina majstora
- Placanje usluge majstora (korisnik ima novčano stanje na profilu, koje može da dopunjava na salteru)
- Ocjenjivanje i ostavljanje recenzija na zavrsene poslove i majstora
- Odgovor na dobijene recenzije, prijava istih

3. Majstor
- Prima obavjestenja i ima uvid u rezervisane termine 
- Prihvatanje rezervisanog posla (nakon čega se korisnik obavještava putem email-a)
- Dodaje i uklanja slobodne termine (radno vrijeme)

4. Radnik
- Pregled i pretraga korisnika po imenu, prezimenu, email-u
- uplata na racun registrovanog korisnika

5. Admin
- Dodavanje majstora
- Pregled i pretraga korisnika po imenu, prezimenu, email-u, roli u sistemu, ocjeni, broju rezervisanih/uradjenih poslova 
- Pregled prijavljenih recenzija nakon čega admin može da blokira korisnika ukoliko uvidi nedolično ponašanje (putem email-a korisnik biva obavješten o blokiranju)

*Sistem pratiti prijavljene recenzije i korisniku se dodjelju negativan bod za neispostavani dogovor, nakon 3 negativna boda, korisnik ce biti blokiran.

## Arhitektura 

GO
- API Gateway: Go
- Mikroservis za korisnike
- Mikroservis za majstore
- Mikroservis za radnike
- Mikroservis za admina
- Mikroservis za rezervacije
- Mikroservis za komentare i ocene 

Pharo
- Mikroservis za email (slanje potvrde o rezervacijama)

Angular
- Klijentska veb aplikacija
