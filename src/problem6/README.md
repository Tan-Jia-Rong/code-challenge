# Market Module

## Overview
The Market Module is responsible for managing the markets on the blockchain. It serves as the backbone of our application, providing users and developers with the capability to create, configure, and operate markets.

It aims to facilitate the exchange of assets within our application.

## Prerequisites

Before installing and configuring the Market Module, make sure you have the following prerequisites in place:

1. **GoLang**: Make sure you have Golang installed on your system. It is required for building and running the Market Module.

2. **Cosmos SDK**: Make sure you are familiar with the usage of Cosmos SDK.

## 1. Quick Start / Installation:

1. Clone the repository
```shell
git clone https://github.com/Switcheo/code-challenge.git
```

### 1.1 Command Summary
| Action| Format| Short Description|
|-------|-------|------------------|
| [Create Market](#121-create-market)| `create-market [market-type] [base] [quote] [base-usd-price] [quote-usd-price] [index-oracle-id] [expiry-time]`| Creates a new market|
| [Show Market](#122-show-market)| `market [name]`| Retrieves detailed information about a specific market given its name|
| [List Market](#123-list-market)| `market`| List all existing markets  |
| [Update Market](#124-update-market)| *Work in progress*| Update the configuration of an existing market|

*Other commands are not included as they are not implemented*
## 1.2 Example Usage

### 1.2.1 Create Market
You can create a new market within your application by passing in desired parameters

This command requires seven arguments
|Arguments|Description|
|-----------|-----------|
|market-type| Type of market|
|base| Base asset|
|quote| Quote asset|
|base-usd-price| Price of 1 unit of base asset in USD|
|quote-usd-price| Price of 1 unit of quote asset in USD|
|index-oracle-id| The oracle service for price data|
|expiry-time| Market's expiry time|

Format: `create-market [market-type] [base] [quote] [base-usd-price] [quote-usd-price] [index-oracle-id] [expiry-time]`


Example Code: `create-market spot BTC ETH 44000 3000 my-oracle-service 1679261256`

Output:
> To be added

### 1.2.2 Show Market

*Work in progress*


### 1.3 Testing
Tests are essential to ensure the functionality and integrity of the module. Hence, examples should be provided on how testing could be done

### 2. Shortcomings and Improvements:

#### Shortcoming 1: Documentation

The Market module lacks a comprehensive and well-structured documentation.
Improvement needs to be made to the documentation by providing clearer explanations, examples, and use cases to help developers understand how to use the module effectively.

Additionally, in each file, it would be good to have a high level overview of what the file is about.

For example, for txMarket.go
```go
package cli

import (
//...
)

// txMarket.go defines the CLI command for market Transactions...


func CmdCreateMarket() *cobra.Command {
    //...
}
```

#### Shortcoming 2: Unused Imports

There are multiple lines of imports statements that are commented-out in some files. If it is not being used, it should be removed to keep the codebase clean.

#### Shortcoming 3: Complicated Expressions
Instead of having two steps in a single line, we can enhance code readability if we split it into two lines.


**Before**
```go
if found := k.HasMarket(ctx, market.Name); found {
    return market, errorsmod.Wrap(types.ErrDuplicateName, market.Name)
}
```
**After**
```go
found := k.HasMarket(ctx, market.Name)
if found {
    return market, errorsmod.Wrap(types.ErrDuplicateName, market.Name)
}
```



### 3. References: Provide links to relevant documentation, frameworks, and dependencies.
- Cosmos SDK Module
    - Folder Structure - https://docs.cosmos.network/main/building-modules/structure
    - Transaction CLI - https://docs.cosmos.network/main/core/cli#transaction-commands
    - MsgServer - https://docs.cosmos.network/v0.47/building-modules/msg-services#state-transition
    - Keeper - https://docs.cosmos.network/v0.47/building-modules/keeper
    - KV Store Reading - https://docs.cosmos.network/v0.46/core/store.html#base-layer-kvstores
- Additional Information
    - KV Store Concept - https://www.influxdata.com/key-value-database/
    - Query CLI - https://docs.cosmos.network/main/core/cli#query-commands
    - Module Migration - https://docs.cosmos.network/main/building-modules/upgrade