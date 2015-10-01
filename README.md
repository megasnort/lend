#Lend

##tl;dr
This is a hobby project and attempt to do two things

- learn go
- make a terminal programma for storing information about who lend a dvd, book, game ... from me

##Concept of usage

Lend an item to someone

```
lend [item] to [human]

```
Register that someone returnes an item

```
lend [human] returned [item]

```
Show an overview of lent items

```
lend lent

```

Create, update or delete a new human

```
lend human [create|update|delete] name_of_human

```

Create, update or delete a new item

```
lend item [create|update|delete] name_of_item

```

Create, update or delete a new category

```
lend cat [create|update|delete] name_of_category

```
Switch between local and online storage

```
lend switch [local|online]

```
Open interactive mode where you don't have to prefix every command with 'lend'

```
lend

```

...

##Features
- Cross platform
- support for parented categories
- store data local
- store data online via an API
- ...