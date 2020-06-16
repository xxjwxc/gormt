# 快捷函数功能
- 说明: 快捷函数只是对 gorm 的辅助功能。目前只支持查询功能
## 目录
 - [_BaseMgr](#_BaseMgr)
	- [_BaseMgr](#_BaseMgr)
	- [SetCtx](#SetCtx)
	- [GetDB](#GetDB)
	- [GetIsRelated](#GetIsRelated)
	- [SetIsRelated](#SetIsRelated)
 - [表逻辑函数](#表逻辑函数)
 	- [简要说明](#简要说明)
    - [逻辑基础类型]](#逻辑基础类型)
    - [已有条件获取方式](#已有条件获取方式)
    - [功能选项方式获取](#功能选项方式获取)
    - [单元素方式获取](#单元素方式获取)
    - [索引方式获取](#索引方式获取)

## _BaseMgr
### OpenRelated : 打开全局预加载
### CloseRelated : 关闭全局预加载
   基础函数。所有管理类型都是继承此函数。此函数提供基础公共函数。
### SetCtx 
    设置 context ，用于设置上下文。目前功能未启用
### GetDB
    获取 gorm.DB 原始链接窜
### UpdateDB
    更新 gorm.DB 原始链接窜
### GetIsRelated
    获取是否查询外键关联
### SetIsRelated
    设置是否查询外键关联

## 表逻辑函数
    表逻辑函数操作关于数据库表相关功能函数:以下使用[xxx]代表逻辑表结构体名
### 简要说明
    查询分为以下几类
### 逻辑基础类型
    `_[xxx]Mgr` : 逻辑表类型
### 已有条件获取方式
    `Get/Gets` : 批量获取(使用 gormt 预先设置的条件可以使用它来获取最终结果)
### 功能选项方式获取

    此功能 用于支持多种条件获取

    `GetByOption/GetByOptions` : 功能选项列表获取
    `With[xxx]` : 功能选项中的参数列表

### 单元素方式获取

    `GetFrom[xxx]` : 元素获取(单个case条件获取)

    `GetBatchFrom[xxx]` : 批量元素获取(单个条件的数组获取) 

### 索引方式获取
    `FetchByPrimaryKey` : 主键获取
    `FetchUniqueBy[xxx]` : 唯一索引方式获取
    `FetchIndexBy[xxx]` : 复合索引获取(返回多个)
    `FetchUniqueIndexBy[xxx]` : 唯一复合索引获取(返回一个)