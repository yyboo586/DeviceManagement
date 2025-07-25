# 设备管理服务 REST API 设计

## 1. 多组织架构支持

### 1.1 组织隔离策略

**策略一：URL路径隔离（推荐）**
```
/api/v1/device-management/orgs/{org_id}/products
/api/v1/device-management/orgs/{org_id}/devices
/api/v1/device-management/orgs/{org_id}/thing-models
```

**策略二：查询参数隔离**
```
/api/v1/device-management/products?org_id={org_id}
/api/v1/device-management/devices?org_id={org_id}
```

**策略三：请求头隔离**
```
X-Org-ID: {org_id}
```

### 1.2 推荐方案：URL路径隔离

**优势：**
- 清晰的资源层级关系
- 便于权限控制
- 符合RESTful设计原则
- 便于API网关路由

## 2. 完整的API接口设计

### 2.1 产品分类管理

```http
# 获取组织下的产品分类树
GET /api/v1/device-management/orgs/{org_id}/product-categories/tree

# 创建产品分类
POST /api/v1/device-management/orgs/{org_id}/product-categories
{
    "pid": 0,
    "name": "传感器类"
}

# 更新产品分类
PUT /api/v1/device-management/orgs/{org_id}/product-categories/{category_id}
{
    "pid": 0,
    "name": "传感器类"
}

# 删除产品分类
DELETE /api/v1/device-management/orgs/{org_id}/product-categories/{category_id}
```

### 2.2 产品管理

```http
# 获取组织下的产品列表
GET /api/v1/device-management/orgs/{org_id}/products?category_id={category_id}&page=1&page_size=10

# 创建产品
POST /api/v1/device-management/orgs/{org_id}/products
{
    "category_id": 2,
    "name": "华为温度传感器",
    "manufacturer": "华为",
    "model": "HW-TEMP-001",
    "description": "工业级温度传感器"
}

# 获取产品详情
GET /api/v1/device-management/orgs/{org_id}/products/{product_id}

# 更新产品
PUT /api/v1/device-management/orgs/{org_id}/products/{product_id}
{
    "category_id": 2,
    "name": "华为温度传感器",
    "manufacturer": "华为",
    "model": "HW-TEMP-001",
    "description": "工业级温度传感器"
}

# 删除产品
DELETE /api/v1/device-management/orgs/{org_id}/products/{product_id}
```

### 2.3 物模型管理

```http
# 获取产品下的物模型列表
GET /api/v1/device-management/orgs/{org_id}/products/{product_id}/thing-models

# 创建物模型（自定义）
POST /api/v1/device-management/orgs/{org_id}/products/{product_id}/thing-models
{
    "name": "华为温度传感器物模型",
    "description": "温度传感器物模型定义",
    "properties": [...],
    "services": [...],
    "events": [...]
}

# 从模板创建物模型
POST /api/v1/device-management/orgs/{org_id}/products/{product_id}/thing-models/from-template
{
    "template_id": 1,
    "name": "华为温度传感器物模型",
    "description": "基于温度传感器模板创建"
}

# 获取物模型详情
GET /api/v1/device-management/orgs/{org_id}/products/{product_id}/thing-models/{thing_model_id}

# 更新物模型
PUT /api/v1/device-management/orgs/{org_id}/products/{product_id}/thing-models/{thing_model_id}
{
    "name": "华为温度传感器物模型",
    "version": "1.1",
    "description": "温度传感器物模型定义",
    "properties": [...],
    "services": [...],
    "events": [...]
}

# 删除物模型
DELETE /api/v1/device-management/orgs/{org_id}/products/{product_id}/thing-models/{thing_model_id}
```

### 2.4 设备管理

```http
# 获取组织下的设备列表
GET /api/v1/device-management/orgs/{org_id}/devices?product_id={product_id}&status={status}&page=1&page_size=10

# 获取产品下的设备列表
GET /api/v1/device-management/orgs/{org_id}/products/{product_id}/devices?status={status}&page=1&page_size=10

# 创建设备
POST /api/v1/device-management/orgs/{org_id}/products/{product_id}/devices
{
    "name": "车间1号温度传感器",
    "device_key": "temp_sensor_001",
    "thing_model_id": 1,
    "thing_model_version": "1.0",
    "device_secret": "secret123",
    "auth_type": 1,
    "location": "车间A区",
    "description": "监控车间温度",
    "tags": {"location": "车间A区", "type": "温度传感器"}
}

# 获取设备详情
GET /api/v1/device-management/orgs/{org_id}/products/{product_id}/devices/{device_id}

# 更新设备
PUT /api/v1/device-management/orgs/{org_id}/products/{product_id}/devices/{device_id}
{
    "name": "车间1号温度传感器",
    "location": "车间A区",
    "description": "监控车间温度",
    "tags": {"location": "车间A区", "type": "温度传感器"}
}

# 删除设备
DELETE /api/v1/device-management/orgs/{org_id}/products/{product_id}/devices/{device_id}

# 设备状态控制
POST /api/v1/device-management/orgs/{org_id}/products/{product_id}/devices/{device_id}/enable
POST /api/v1/device-management/orgs/{org_id}/products/{product_id}/devices/{device_id}/disable
```

### 2.5 物模型模板管理

```http
# 获取系统物模型模板列表
GET /api/v1/device-management/thing-model-templates?category_id={category_id}

# 获取模板详情
GET /api/v1/device-management/thing-model-templates/{template_id}

# 创建自定义模板
POST /api/v1/device-management/orgs/{org_id}/thing-model-templates
{
    "category_id": 2,
    "name": "自定义温度传感器模板",
    "description": "适用于特定场景的温度传感器",
    "properties": [...],
    "services": [...],
    "events": [...]
}
```

## 3. 请求/响应格式

### 3.1 标准请求头

```http
Content-Type: application/json
Authorization: Bearer {token}
X-Request-ID: {request_id}
```

### 3.2 标准响应格式

```json
{
    "code": 200,
    "message": "success",
    "data": {
        "id": 1,
        "name": "华为温度传感器",
        "created_at": "2024-01-01T00:00:00Z"
    },
    "request_id": "req_123456"
}
```

### 3.3 分页响应格式

```json
{
    "code": 200,
    "message": "success",
    "data": {
        "list": [...],
        "pagination": {
            "page": 1,
            "page_size": 10,
            "total": 100,
            "total_pages": 10
        }
    }
}
```

## 4. 错误处理

### 4.1 标准错误响应

```json
{
    "code": 400,
    "message": "参数验证失败",
    "errors": [
        {
            "field": "name",
            "message": "产品名称不能为空"
        }
    ],
    "request_id": "req_123456"
}
```

### 4.2 常见错误码

- `400` - 请求参数错误
- `401` - 未授权
- `403` - 权限不足
- `404` - 资源不存在
- `409` - 资源冲突
- `422` - 业务逻辑错误
- `500` - 服务器内部错误

## 5. 权限控制

### 5.1 组织级权限

- 用户只能访问自己组织的资源
- 通过URL路径中的`org_id`进行隔离
- 在中间件中验证用户是否有权限访问该组织

### 5.2 资源级权限

- 产品分类：组织内可见
- 产品：组织内可见
- 物模型：产品内可见
- 设备：产品内可见

## 6. API版本控制

### 6.1 版本策略

- 使用URL路径版本：`/api/v1/`
- 向后兼容原则
- 废弃的API通过文档标注

### 6.2 版本升级

- 新版本保持向后兼容
- 提供迁移指南
- 设置合理的废弃期 