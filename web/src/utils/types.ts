import { NetworkType } from "./consts";

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
  type: NetworkType;
  local_ip: string;
  local_port: number;
  remote_port: number;
}
