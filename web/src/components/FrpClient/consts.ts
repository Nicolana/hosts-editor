export const frpTypes = [
    {
        label: "tcp",
        value: "tcp",
    },
    {
        label: "udp",
        value: "udp",
    },
    {
        label: "xtcp",
        value: "xtcp",
    },
    {
        label: "tcpmux",
        value: "tcpmux",
    },
    {
        label: "stcp",
        value: "stcp",
    },
    {
        label: "https",
        value: "https",
    },
    {
        label: "http",
        value: "http",
    },
];

// 定义FRP 配置项表单类型
export enum FrpFormItemType {
    text = 'text',
    select = 'select',
    radio = 'radio',
    number = 'number',
}

export const FrpFormItemEnum = {
    [FrpFormItemType.text]: {
        text: "文本",
    },
    [FrpFormItemType.radio]: {
        text: "单选",
    },
    [FrpFormItemType.select]: {
        text: "下拉选择",
    },
}

export interface FrpConfigItemType {
    name: string
    value: any
    type: FrpFormItemType
}

export const FrpClientConfig = [
    {
        name: 'type',
        type: FrpFormItemType.text,
        default: '',
    },
    {
        name: 'local_ip',
        type: FrpFormItemType.text,
        default: '127.0.0.1',
    },
    {
        name: 'local_port',
        type: FrpFormItemType.number,
        default: 8000,
    },
    {
        name: 'local_port',
        type: FrpFormItemType.number,
        default: 8000,
    },
    {
        name: 'sk',
        type: FrpFormItemType.text,
        default: '',
    },
]