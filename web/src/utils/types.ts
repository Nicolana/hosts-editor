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
