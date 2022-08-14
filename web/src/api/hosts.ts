import { request } from "../utils/http";

export const getHosts = (params: any) => {
  return request.get("/hosts/list", { params });
};

export const deletHosts = (data: any) => {
  return request.post("/hosts/delete", data);
};

export const updateHosts = (data: any) => {
  return request.post("/api/update", data);
};

export const addHosts = (data: any) => {
  return request.post("/api/add", data);
};
