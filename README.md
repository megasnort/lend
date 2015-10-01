#Lend

##tl;dr
This is an ongoing hobby project and an attempt to do two things:

- Learn [go](https://golang.org)
- Make a terminal program for keeping track about to whom I lent a cd, dvd, book, game ...

##Concept of usage

Lend an item to someone

```lend [item] to [human]```

Register that someone returns an item

```lend [human] returned [item]```

Show an overview of lent items

```lend lent```

Create, update or delete a new human

```lend human [create|update|delete] name_of_human```

Create, update or delete a new item

```lend item [create|update|delete] name_of_item```

Create, update or delete a new category

```lend cat [create|update|delete] name_of_category```

Switch between local and online storage

```lend switch [local|online]```

Open interactive mode where you don't have to prefix every command above with 'lend'

```lend```

...


##Features

- cross platform
- support for parented categories
- store data local
- store data online via an API
- auto complete
- ...
- 