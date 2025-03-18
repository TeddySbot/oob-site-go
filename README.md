# TeddysBot OOB Site Go

Bienvenue dans le projet **TeddysBot OOB Site Go** ! Ce site est en cours de développement pour un serveur Garry's Mod (GMod) et s'inspire de l'univers d'**Owari no Seraph**. 

## Structure du Répertoire

Voici la structure des répertoires de ce projet :


Directory structure:
└── teddysbot-oob-site-go/
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── backend/
    │   └── handler/
    │       └── Handler.go
    └── frontend/
        ├── js/
        │   └── index.js
        ├── style/
        │   └── general/
        │       ├── boutique.css
        │       ├── index.css
        │       ├── lien.css
        │       ├── lore.css
        │       └── reglement.css
        └── template/
            ├── ad/
            │   ├── ad.html
            │   └── reglement-ad.html
            ├── general/
            │   ├── boutique.html
            │   ├── index.html
            │   ├── lien.html
            │   ├── lore.html
            │   └── reglement.html
            └── vampire/
                ├── reglement-vampire.html
                └── vampire.html

# Exécution

<h9> Pour démarrer le serveur, exécutez les commandes suivantes :</h9>

**Vérifier que vous avez bien git et golang de télécharger pour pouvoir exécuter le programme.**

```powershell
git clone https://github.com/TeddySbot/oob-site-go.git
```


```powershell
go run main.go
```