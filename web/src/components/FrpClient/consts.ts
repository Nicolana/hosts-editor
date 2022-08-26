export enum NetworkType {
  TCP = "tcp",
  UDP = "udp",
  XTCP = "xtcp",
  TCPMUX = "tcpmux",
  STCP = "stcp",
  HTTPS = "https",
  HTTP = "http",
}

// 定义FRP 配置项表单类型
export enum FrpFormItemType {
  text = "text",
  select = "select",
  radio = "radio",
  number = "number",
  enum = "enum",
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
};

export interface FrpConfigItemType {
  name: string;
  key: string;
  value?: any;
  type: FrpFormItemType;
  default?: any;
  valueEnum?: { [keyProp: string]: any };
  min?: number;
  max?: number;
}

export const FrpClientConfig: FrpConfigItemType[] = [
  {
    name: "类型",
    key: "type",
    type: FrpFormItemType.enum,
    valueEnum: {
      [NetworkType.TCP]: {
        text: "TCP",
      },
      [NetworkType.UDP]: {
        text: "UDP",
      },
      [NetworkType.XTCP]: {
        text: "XTCP",
      },
      [NetworkType.TCPMUX]: {
        text: "TCPMUX",
      },
      [NetworkType.STCP]: {
        text: "STCP",
      },
      [NetworkType.HTTPS]: {
        text: "HTTPS",
      },
      [NetworkType.HTTP]: {
        text: "HTTP",
      },
    },
  },
  {
    name: "内网地址",
    key: "local_ip",
    type: FrpFormItemType.text,
    value: "127.0.0.1",
  },
  {
    name: "内网端口",
    key: "local_port",
    type: FrpFormItemType.number,
    value: 8000,
    min: 0,
    max: 65535,
  },
  {
    name: "服务器端口",
    key: "remote_port",
    type: FrpFormItemType.number,
    value: 8000,
    min: 0,
    max: 65535,
  },
  {
    name: "sk",
    key: "sk",
    type: FrpFormItemType.text,
    value: "",
  },
  {
    name: "带宽限制",
    key: "bandwidth_limit",
    type: FrpFormItemType.text,
    value: "",
  },
  {
    name: "加密传输",
    key: "use_encryption",
    type: FrpFormItemType.radio,
    value: false,
  },
];

export const GetFormItemByKey = (key: string) => {
  return FrpClientConfig.find((item) => item.key === key);
};
