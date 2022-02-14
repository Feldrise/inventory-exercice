# GraphQL exercise

This is a simple GraphQL exercise in Golang with gqlgen

## Exemple clients
There is serval exemple of clients using this API

| Language | Framework | GitHub |
| -------- | --------- | ------ |
| Dart     | Flutter & [graphql_flutter](https://pub.dev/packages/graphql_flutter) | [Link](https://github.com/Feldrise/Flutter-GraphQL-Inventory) |
| TypeScript | NextJS & [Apollo Client](https://www.apollographql.com/) | *Comming Soon* |

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
    email: "me1@me.com",
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

mutation refreshToken {
  refreshToken(input: {
    token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQ2NjA4NzcsImlkIjoiNjIwNGY5YzM4NWEyZmE4ZWU5NzQzOTBmIn0.bgIgbylrgJWJ_mQydwgtI8WyeUq5TM9n1GeTlepo8ik"
  })
}

#######################
## INVENTORIES STUFF ##
#######################

## Inventories

# Mutations

mutation createInventory {
  createInventory(input: {
    name: "Inventory 2",
    description: "This is the second inventory created"
  }) {
    name,
    user {
      email
    }
  }
}

mutation updateInventory {
  updateInventory(id: "620639c66d395c3f60a8463f", changes: {
    name: "Updated Inventory"
  }) {
    name
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
    # To learn more about pagination : https://www.apollographql.com/blog/graphql/pagination/understanding-pagination-rest-graphql-and-relay/
    items(first: 5) {
      edges {
        node {
          name,
          quantity
        }
      },
      pageInfo {
        startCursor,
        endCursor
      }
    }
  }
}

query getInventory {
  inventory(id: "6208be109403b4405ad4c54d") {
    name,
    # To learn more about pagination : https://www.apollographql.com/blog/graphql/pagination/understanding-pagination-rest-graphql-and-relay/
    items(first: 5, after: "NjIwYTJjNTE3YjgxOTliOTgzMWRiNDlj") {
      edges {
        node {
          name,
          quantity
        }
      },
      pageInfo {
        startCursor,
        endCursor
      }
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

mutation updateItem {
  updateInventoryItem(id: "62064709a6e21068b8d11dc3", changes: {
    name: "Updated Item 1"
    quantity: 3
  }) {
    name,
    quantity
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