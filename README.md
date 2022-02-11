# GraphQL exercise

This is a simple GraphQL exercise in Golang with gqlgen

## Exemple queries
Here you have many exemple that can be use to play around with the API

```graphql
#################
## USERS STUFF ##
#################

mutation createUser1 {
  createUser(input: {
    email: "admin@me.com",
    password: "dE8bdTUE"
  }) {
    id,
    email
  }
}


mutation createUser2 {
  createUser(input: {
    email: "user@me.com",
    password: "dE8bdTUE"
  }) {
    id,
    email
  }
}

# Token : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQ2NjA4NzcsImlkIjoiNjIwNGY5YzM4NWEyZmE4ZWU5NzQzOTBmIn0.bgIgbylrgJWJ_mQydwgtI8WyeUq5TM9n1GeTlepo8ik
mutation login1 {
  login(input: {
    email: "admin@me.com",
    password: "dE8bdTUE"
  })
}

# Token : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQ2NjI0ODcsImlkIjoiNjIwNGZhZTJjNTNmNzFkNzZjZmEwOWQ5In0.ifsp6O7XNlkHiy7qNdgEV_zQKnicJqWC0J2xieO1ojo
mutation login2 {
  login(input: {
    email: "user@me.com",
    password: "dE8bdTUE"
  })
}

mutation failedLogin1 {
  login(input: {
    email: "admin@me.com",
    password: "dE8bdUE"
  })
}

mutation failedLogin2 {
  login(input: {
    email: "adm@me.com",
    password: "dE8bdTUE"
  })
}

#######################
## INVENTORIES STUFF ##
#######################

## Inventories

# Mutations

mutation createInventory {
  createInventory(input: {
    name: "My first inventory"
  }) {
    name,
    user {
      email
    }
  }
}

# Queries

query getInventories {
  inventories {
    id
    name,
    user {
      email
    },
    items {
      name,
      quantity
    }
  }
}

query getInventory {
  inventory(id: "620639c66d395c3f60a8463f") {
    name,
    items {
      name,
      quantity
    }
  }
}

## Items

# Mutation

mutation createItem {
  createInventoryItem(input: {
    inventoryID: "620639c66d395c3f60a8463f",
    name: "Item 3",
    quantity: 5 # optional
  }) {
    name
  }
}

# Query

query getItem {
  inventoryItem(id: "62064709a6e21068b8d11dc3") {
    name,
    quantity
  }
}
```