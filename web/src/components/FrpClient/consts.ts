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
  switch = "switch",
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

export const healthCheckEnum = {
  [NetworkType.TCP]: {
    text: "TCP",
  },
  [NetworkType.HTTP]: {
    text: "HTTP",
  },
};

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
    type: FrpFormItemType.switch,
    value: false,
  },
  {
    name: "压缩传输",
    key: "use_compression",
    type: FrpFormItemType.switch,
    value: false,
  },
  // # frps will load balancing connections for proxies in same group
  {
    name: "group",
    key: "group",
    type: FrpFormItemType.text,
    value: "",
  },
  // # group should have same group key
  {
    name: "group_key",
    key: "group_key",
    type: FrpFormItemType.text,
    value: "",
  },
  // # enable health check for the backend service, it support 'tcp' and 'http' now
  // # frpc will connect local service's port to detect it's healthy status
  {
    name: "health_check_type",
    key: "health_check_type",
    type: FrpFormItemType.enum,
    valueEnum: healthCheckEnum,
  },
  // # health check connection timeout，单位秒
  {
    name: "health_check_timeout_s",
    key: "health_check_timeout_s",
    type: FrpFormItemType.number,
    value: 3,
  },
  // # if continuous failed in 3 times, the proxy will be removed from frps
  {
    name: "health_check_max_failed",
    key: "health_check_max_failed",
    type: FrpFormItemType.number,
    value: 3,
  },
  //  # every 10 seconds will do a health check
  {
    name: "health_check_interval_s",
    key: "health_check_interval_s",
    type: FrpFormItemType.number,
    value: 10,
  },
  // # http username and password are safety certification for http protocol
  // # if not set, you can access this custom_domains without certification
  {
    name: "http_user",
    key: "http_user",
    type: FrpFormItemType.text,
    value: "admin",
  },
  {
    name: "http_pwd",
    key: "http_pwd",
    type: FrpFormItemType.text,
    value: "admin",
  },
  // # if domain for frps is frps.com, then you can access [web01] proxy by URL http://web01.frps.com
  {
    name: "subdomain",
    key: "subdomain",
    type: FrpFormItemType.text,
    value: "web01",
  },
  {
    name: "custom_domains",
    key: "custom_domains",
    type: FrpFormItemType.text,
    value: "web01.yourdomain.com",
  },
  // # locations is only available for http type
  {
    name: "locations",
    key: "locations",
    type: FrpFormItemType.text,
    value: "/,/pic",
  },
  {
    name: "host_header_rewrite",
    key: "host_header_rewrite",
    type: FrpFormItemType.text,
    value: "example.com",
  },
  // # if plugin is defined, local_ip and local_port is useless
  // # plugin will handle connections got from frps
  {
    name: "plugin",
    key: "plugin",
    type: FrpFormItemType.text,
    value: "unix_domain_socket",
  },
  // # params with prefix "plugin_" that plugin needed
  {
    name: "plugin_unix_path",
    key: "plugin_unix_path",
    type: FrpFormItemType.text,
    value: "/var/run/docker.sock",
  },
  {
    name: "bind_addr",
    key: "bind_addr",
    type: FrpFormItemType.text,
    value: "127.0.0.1",
  },
  {
    name: "bind_port",
    key: "bind_port",
    type: FrpFormItemType.number,
    value: 9001,
  },
  {
    name: "server_name",
    key: "server_name",
    type: FrpFormItemType.text,
    value: "secret_tcp",
  },
  {
    name: "role",
    key: "role",
    type: FrpFormItemType.text,
    value: "visitor",
  },
];

export const FrpClientConfigKeyMap = {} as { [prop: string]: any };
for (const index in FrpClientConfig) {
  const item = FrpClientConfig[index];
  FrpClientConfigKeyMap[item.key] = index;
}

export const GetFormItemByKey = (key: string) => {
  const index =
    FrpClientConfigKeyMap[key] > -1 ? FrpClientConfigKeyMap[key] : 0;
  return { ...FrpClientConfig[index] };
};
