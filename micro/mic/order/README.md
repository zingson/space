# 订单服务(order)

机构用户订单、会员用户订单

不同店铺商品自动拆单
虚拟与实物商品，自动拆单

查询产品功能字段，调用对应处理逻辑，如虚拟产品自动发货


## 数据结构

会员订单信息(ms_order)
订单信息（订单号、订单金额、订单、订单明细、支付状态、发货状态、支付超时时间、支付完成时间、订单创建时间、下单会员ID、商城ID）




## 接口

- 创建订单（用户ID、订单ID、订单金额、订单支付状态（1=待支付 2=支付 3=）、订单明细（产品ID、产品名称、产品价格、产品数量、产品图片、））
- 订单列表
- 订单详情
