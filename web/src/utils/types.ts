import { FrpConfigItemType } from "@/components/FrpClient/consts";

export interface PayloadTypes {
  row?: number;
  hosts: string;
  ip: string;
  comments?: string;
}

export interface PaginationType {
  page: number;
  size: number;
  total: number;
  search?: string;
}

export interface FrpPayloadTypes {
  name?: string;
  frpConfigs: FrpConfigItemType[];
}

export interface FrpServerTypes {
  server_addr: string;
  server_port: number;
  token: string;
}
