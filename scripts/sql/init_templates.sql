-- Product Category Initialization Data
INSERT INTO `t_product_category` ( `id`, `pid`, `org_id`, `name`) VALUES 
(
    1, 0, '1', 'Industrial Equipment'
);
INSERT INTO `t_product_category` ( `id`, `pid`, `org_id`, `name`) VALUES 
(
    2, 1, '1', 'Sensor Class'
);
INSERT INTO `t_product_category` ( `id`, `pid`, `org_id`, `name`) VALUES 
(
    3, 1, '1', 'Controller Class'
);

-- Thing Model Template Initialization Data
-- Temperature Sensor Template
INSERT INTO `t_thing_model_template` 
( `id`, `name`, `description`, `properties`, `services`, `events`, `is_system`) VALUES 
(
    1, 'Temperature Sensor Template', 'Applicable to industrial temperature sensor devices',
    '[
        {
            "id": "temperature",
            "name": "Temperature",
            "data_type": "float",
            "unit": "°C",
            "min": -50,
            "max": 100,
            "required": true,
            "description": "Current temperature value"
        },
        {
            "id": "humidity",
            "name": "Humidity",
            "data_type": "float",
            "unit": "%",
            "min": 0,
            "max": 100,
            "required": false,
            "description": "Current humidity value"
        },
        {
            "id": "battery",
            "name": "Battery Level",
            "data_type": "int",
            "unit": "%",
            "min": 0,
            "max": 100,
            "required": false,
            "description": "Battery remaining level"
        }
    ]',
    '[
        {
            "id": "setThreshold",
            "name": "Set Threshold",
            "description": "Set temperature alarm threshold",
            "input": [
                {
                    "id": "minTemp",
                    "name": "Minimum Temperature",
                    "data_type": "float",
                    "required": true,
                    "description": "Minimum temperature threshold"
                },
                {
                    "id": "maxTemp",
                    "name": "Maximum Temperature",
                    "data_type": "float",
                    "required": true,
                    "description": "Maximum temperature threshold"
                }
            ],
            "output": [
                {
                    "id": "result",
                    "name": "Set Result",
                    "data_type": "bool",
                    "description": "Whether the setting is successful"
                }
            ]
        },
        {
            "id": "getConfig",
            "name": "Get Configuration",
            "description": "Get device configuration information",
            "input": [],
            "output": [
                {
                    "id": "reportInterval",
                    "name": "Report Interval",
                    "data_type": "int",
                    "description": "Data reporting interval (seconds)"
                },
                {
                    "id": "threshold",
                    "name": "Threshold Configuration",
                    "data_type": "object",
                    "description": "Temperature threshold configuration"
                }
            ]
        }
    ]',
    '[
        {
            "id": "temperatureAlert",
            "name": "Temperature Alert",
            "description": "Trigger alarm when temperature exceeds threshold",
            "output": [
                {
                    "id": "temperature",
                    "name": "Current Temperature",
                    "data_type": "float",
                    "description": "Temperature value when alarm is triggered"
                },
                {
                    "id": "threshold",
                    "name": "Trigger Threshold",
                    "data_type": "float",
                    "description": "Threshold that triggered the alarm"
                },
                {
                    "id": "alertType",
                    "name": "Alert Type",
                    "data_type": "string",
                    "description": "High temperature alarm or low temperature alarm"
                }
            ]
        }
    ]',
    0
);

-- Humidity Sensor Template
INSERT INTO `t_thing_model_template` 
( `id`, `name`, `description`, `properties`, `services`, `events`, `is_system`) VALUES 
(
    2, 'Humidity Sensor Template', 'Applicable to industrial humidity sensor devices',
    '[
        {
            "id": "humidity",
            "name": "Humidity",
            "data_type": "float",
            "unit": "%",
            "min": 0,
            "max": 100,
            "required": true,
            "description": "Current humidity value"
        },
        {
            "id": "temperature",
            "name": "Temperature",
            "data_type": "float",
            "unit": "°C",
            "min": -50,
            "max": 100,
            "required": false,
            "description": "Current temperature value"
        }
    ]',
    '[
        {
            "id": "setHumidityThreshold",
            "name": "Set Humidity Threshold",
            "description": "Set humidity alarm threshold",
            "input": [
                {
                    "id": "minHumidity",
                    "name": "Minimum Humidity",
                    "data_type": "float",
                    "required": true,
                    "description": "Minimum humidity threshold"
                },
                {
                    "id": "maxHumidity",
                    "name": "Maximum Humidity",
                    "data_type": "float",
                    "required": true,
                    "description": "Maximum humidity threshold"
                }
            ],
            "output": [
                {
                    "id": "result",
                    "name": "Set Result",
                    "data_type": "bool",
                    "description": "Whether the setting is successful"
                }
            ]
        }
    ]',
    '[
        {
            "id": "humidityAlert",
            "name": "Humidity Alert",
            "description": "Trigger alarm when humidity exceeds threshold",
            "output": [
                {
                    "id": "humidity",
                    "name": "Current Humidity",
                    "data_type": "float",
                    "description": "Humidity value when alarm is triggered"
                },
                {
                    "id": "threshold",
                    "name": "Trigger Threshold",
                    "data_type": "float",
                    "description": "Threshold that triggered the alarm"
                }
            ]
        }
        ]',
    0
);

-- PLC Controller Template
INSERT INTO `t_thing_model_template` ( `id`, `name`, `description`, `properties`, `services`, `events`, `is_system`) VALUES 
(3, 'PLC Controller Template', 'Applicable to industrial PLC controller devices',
    '[
        {
            "id": "status",
            "name": "Running Status",
            "data_type": "int",
            "required": true,
            "description": "PLC running status (0: Stop 1: Running 2: Fault)"
        },
        {
            "id": "mode",
            "name": "Running Mode",
            "data_type": "int",
            "required": true,
            "description": "Running mode (0: Manual 1: Automatic)"
        },
        {
            "id": "temperature",
            "name": "CPU Temperature",
            "data_type": "float",
            "unit": "°C",
            "required": false,
            "description": "CPU temperature"
        }
    ]',
    '[
        {
            "id": "start",
            "name": "Start",
            "description": "Start PLC controller",
            "input": [],
            "output": [
                {
                    "id": "result",
                    "name": "Start Result",
                    "data_type": "bool",
                    "description": "Whether the start is successful"
                }
            ]
        },
        {
            "id": "stop",
            "name": "Stop",
            "description": "Stop PLC controller",
            "input": [],
            "output": [
                {
                    "id": "result",
                    "name": "Stop Result",
                    "data_type": "bool",
                    "description": "Whether the stop is successful"
                }
            ]
        },
        {
            "id": "setMode",
            "name": "Set Mode",
            "description": "Set running mode",
            "input": [
                {
                    "id": "mode",
                    "name": "Running Mode",
                    "data_type": "int",
                    "required": true,
                    "description": "Running mode (0: Manual 1: Automatic)"
                }
            ],
            "output": [
                {
                    "id": "result",
                    "name": "Set Result",
                    "data_type": "bool",
                    "description": "Whether the setting is successful"
                }
            ]
        }
    ]',
    '[
        {
            "id": "statusChange",
            "name": "Status Change",
            "description": "Triggered when PLC status changes",
            "output": [
                {
                    "id": "oldStatus",
                    "name": "Previous Status",
                    "data_type": "int",
                    "description": "Status before change"
                },
                {
                    "id": "newStatus",
                    "name": "New Status",
                    "data_type": "int",
                    "description": "Status after change"
                },
                {
                    "id": "timestamp",
                    "name": "Timestamp",
                    "data_type": "string",
                    "description": "Time when status changed"
                }
            ]
        },
        {
            "id": "fault",
            "name": "Fault Alert",
            "description": "Triggered when PLC fault occurs",
            "output": [
                {
                    "id": "faultCode",
                    "name": "Fault Code",
                    "data_type": "int",
                    "description": "Fault code"
                },
                {
                    "id": "faultMessage",
                    "name": "Fault Message",
                    "data_type": "string",
                    "description": "Detailed fault information"
                }
            ]
        }
    ]',
    0
);