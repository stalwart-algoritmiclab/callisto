## Unreleased
- ([\#610](https://github.com/forbole/callisto/pull/610)) Add support for gov `v1` proposals
- ([\#652](https://github.com/forbole/callisto/pull/652)) Update gov module parsing 
- ([\#702](https://github.com/forbole/callisto/pull/702)) Add `message_type` module and store msg types inside `message_type` table, add `messages_by_type` function to allow to query messages by their types
- ([\#704](https://github.com/forbole/callisto/pull/704)) Update `callisto` name to `Callisto`

## Version v4.0.0
## Notes
This version is thought to be used with Cosmos SDK `v0.47.x`.

### Changes

#### Parse Command
- ([\#492](https://github.com/forbole/callisto/pull/492)) Add parse command for periodic tasks: `x/bank` total supply, `x/distribution` community pool, `x/mint` inflation, `pricefeed` token price and price history, `x/staking` staking pool

#### Upgrade Module
- ([\#467](https://github.com/forbole/callisto/pull/467)) Store software upgrade plan and refresh data at upgrade height

#### Staking Module
- ([\#443](https://github.com/forbole/callisto/pull/443)) Remove tombstone status from staking module(already stored in slashing module)
- ([\#455](https://github.com/forbole/callisto/pull/455)) Added `unbonding_tokens` and `staked_not_bonded_tokens` values to staking pool table
- ([\#536](https://github.com/forbole/callisto/pull/536)) Fix `PoolSnapshot` tokens type from  `sdk.Int` to `sdkmath.Int`

#### Gov Module
- ([\#461](https://github.com/forbole/callisto/pull/461)) Parse `x/gov` genesis with `genesisDoc.InitialHeight` instead of the hard-coded height 1
- ([\#465](https://github.com/forbole/callisto/pull/465)) Get open proposal ids in deposit or voting period by block time instead of current time
- ([\#489](https://github.com/forbole/callisto/pull/489)) Remove block height foreign key from proposal_vote and proposal_deposit tables and add column timestamp
- ([\#499](https://github.com/forbole/callisto/pull/499)) Check if proposal has passed voting end time before marking it invalid
- ([\#523](https://github.com/forbole/callisto/pull/523)) Update proposal snapshots handling on block
- ([\#681](https://github.com/forbole/callisto/pull/681)) Handle proposal status change from deposit to voting
- 
#### Daily refetch
- ([\#454](https://github.com/forbole/callisto/pull/454)) Added `daily refetch` module to refetch missing blocks every day

#### Hasura
- ([\#473](https://github.com/forbole/callisto/pull/473)) Improved Hasura permissions
- ([\#491](https://github.com/forbole/callisto/pull/491)) Add host address to Hasura actions

### Dependencies
- ([\#542](https://github.com/forbole/callisto/pull/542)) Updated Juno to `v5.1.0`


## Version v3.2.0
### Changes
#### Mint module
- ([\#432](https://github.com/forbole/callisto/pull/432)) Update inflation rate when mint param change proposal is passed

#### Gov module
- ([\#401](https://github.com/forbole/callisto/pull/401)) Update the proposal status to the latest in `callisto parse gov proposal [id]` command
- ([\#430](https://github.com/forbole/callisto/pull/430)) Update the proposals that have invalid status but can still be in voting or deposit periods 

### Dependencies
- ([\#440](https://github.com/forbole/callisto/pull/440)) Updated Juno to `v3.3.0`

## Version v3.1.0
### Dependencies
- Updated Juno to `v3.2.0`

### Changes 
#### Hasura
- ([\#395](https://github.com/forbole/callisto/pull/395)) Remove time label from Hasura Prometheus monitoring

#### Bank module
- ([\#410](https://github.com/forbole/callisto/pull/410)) Change total supply query from only 1 page to all pages

## Version v3.0.1
### Dependencies
- Updated Juno to `v3.1.1`

## Version v3.0.0
### Notes
This version introduces breaking changes to `transaction` and `message` PostgreSQL tables. It implements PostgreSQL table partitioning to fix slow data retrieval from database that stores large amount of transactions and messages. Read more details about [migrating to v3.0.0](https://docs.bigdipper.live/cosmos-based/parser/migrations/v2.0.0)

### New features 
#### CLI
- ([\#356](https://github.com/forbole/callisto/pull/356)) Implemented `migrate` command to perform easy migration to higher callisto versions
- ([\#356](https://github.com/forbole/callisto/pull/356)) Updated `parse-genesis` command to parse genesis file without accessing the node

#### Database
- ([\#356](https://github.com/forbole/callisto/pull/356)) Added PostgreSQL table partition to `transaction` and `message` table
- ([\#356](https://github.com/forbole/callisto/pull/356)) Created new `messages_by_address` function

### Changes 
#### Hasura
- ([\#377](https://github.com/forbole/callisto/pull/377)) Updated Hasura metadata
- ([\#381](https://github.com/forbole/callisto/pull/381)) Hasura actions are now a module 

### Dependencies
- ([\#356](https://github.com/forbole/callisto/pull/356)) Updated Juno to `v3.0.0`

## Version v2.0.0
### Notes
This version introduces breaking changes to certain address-specific data that is no longer periodically parsed from the node and stored in the database. Instead, the data is now obtained directly from the node when needed using Hasura Actions. Read more details about [migrating to v2.0.0](https://docs.bigdipper.live/cosmos-based/parser/migrations/v2.0.0)

### New features
#### CLI
- ([\#257](https://github.com/forbole/callisto/pull/257)) Added `parse-genesis` command to parse the genesis file
- ([\#228](https://github.com/forbole/callisto/pull/228)) ([\#248](https://github.com/forbole/callisto/pull/248)) Added `fix` command:
  - `auth`: fix vesting accounts details
  - `blocks`: fix missing blocks and transactions from given start height
  - `gov`: fix proposal with given proposal ID  
  - `staking`: fix validators info at the latest height  

#### Hasura Actions
- ([\#329](https://github.com/forbole/callisto/pull/329)) Implemented Hasura Actions service to replace periodic queries. If you are using GraphQL queries on your application, you should updated the old queries to use the below new actions instead.  
  Here's a list of data acquired through Hasura Actions:
    - Of a certain address/delegator:
      - Account balance (`action_account_balance`)
      - Delegation rewards (`action_delegation_reward`)
      - Delegator withdraw address (`action_delegator_withdraw_address`)
      - Delegations (`action_delegation`)
      - Total delegations amount (`action_delegation_total`)
      - Unbonding delegations (`action_unbonding_delegation`)
      - Total unbonding delegations amount (`action_unbonding_delegation_total`)
      - Redelegations (`action_redelegation`)
    - Of a certain validator:
      - Commission amount (`action_validator_commission_amount`)
      - Delegations to this validator (`action_validator_delegations`)
      - Redelegations from this validator (`action_validator_redelegations_from`)
      - Unbonding delegations (`action_validator_unbonding_delegations`)
- ([\#352](https://github.com/forbole/callisto/pull/352)) Added prometheus monitoring to hasura actions

#### Local node support
- Added the support for `node.type = "local"` for parsing a static local node without the usage gRPC queries: [config reference](https://docs.bigdipper.live/cosmos-based/parser/config/config#node).

#### Modules
- ([\#232](https://github.com/forbole/callisto/pull/232)) Updated the `x/auth` module support to handle and store `vesting accounts` and `vesting periods` inside the database. 
- ([\#276](https://github.com/forbole/callisto/pull/276)) Added the support for the `x/feegrant` module (v0.44.x)

### Changes 

#### CLI
- ([\#351](https://github.com/forbole/callisto/pull/351)) Fixed version display for `callisto version` cmd 

#### Database
- ([\#300](https://github.com/forbole/callisto/pull/300)) Changed `bonded_tokens` and `not_bonded_tokens` type inside `staking_pool` table  to `TEXT` to avoid value overflow
- ([\#275](https://github.com/forbole/callisto/pull/275)) Added `tombstoned` column inside `validator_status` table
- ([\#232](https://github.com/forbole/callisto/pull/232)) Added `vesting_account` and `vesting_period` table
- ([\#276](https://github.com/forbole/callisto/pull/276)) Added `fee_grant_allowance` table (v0.44.x)

#### Modules
- ([\#353](https://github.com/forbole/callisto/pull/353)) Removed the support for the `history` module
