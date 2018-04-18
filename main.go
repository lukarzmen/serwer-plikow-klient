package main

func main() {
	klient := Klient{
		NazwaKlienta:                "lmedyk",                //os.Getenv("NAZWA_KLIENTA"),
		SciezkaDoFolderuUzytkownika: "/Users/lmedyk/Desktop", //os.Getenv("SCIEZKA"),
		SerwerPlikowURL:             "http://localhost:5050",
	}
	klient.Sledz()
}
