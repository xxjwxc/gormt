# func export readme
- Note: the shortcut function is only an auxiliary function of Gorm. Currently only query function is supported
## Catalog
 - [_BaseMgr](#_BaseMgr)
	- [_BaseMgr](#_BaseMgr)
	- [SetCtx](#SetCtx)
	- [GetDB](#GetDB)
	- [GetIsRelated](#GetIsRelated)
	- [SetIsRelated](#SetIsRelated)
 - [Table logic function](#Table-logic-function)
 	- [Brief description](#Brief-description)
    - [Basic types of logic]](#Basic-types-of-logic)
    - [Access to existing conditions](#Access-to-existing-conditions)
    - [Access to function options](#Access-to-function-options)
    - [Single element access](#Single-element-access)
    - [Index access](#Index-access)

## _BaseMgr

### OpenRelated : open global related
### CloseRelated : close global related

   Basic function. All management types inherit this function. This function provides the underlying common function.
### SetCtx 
    Set context, which is used to set context. The current function is not enabled
### GetDB
    Get gorm.db original link
### UpdateDB
    Update gorm.db original link
### GetIsRelated
    Get whether to query foreign key Association
### SetIsRelated
    Set whether to query foreign key Association

## Table logic function
    Table logical function operation about database table related function: use [XXX] to represent logical table structure name
### Brief description
    Queries are divided into the following categories
### Basic types of logic
    `_[xxx]Mgr` : Logical table type
### Access to existing conditions
    `Get/Gets` : Batch get (you can use gormt to get the final result using its preset conditions)

### Access to function options

    This function is used to support multiple condition acquisition

    `GetByOption/GetByOptions` : Get function option list
    `With[xxx]` : Parameter list in function options

### Single element access

    `GetFrom[xxx]` : Element acquisition (single case conditional acquisition)

    `GetBatchFrom[xxx]` : Batch element acquisition (array acquisition of a single condition) 

### Index access

    `FetchByPrimaryKey` : Primary key acquisition
    `FetchUniqueBy[xxx]` : Get by unique index
    `FetchIndexBy[xxx]` : Composite index fetch (multiple returned)
    `FetchUniqueIndexBy[xxx]` : Unique composite index fetch (return one)