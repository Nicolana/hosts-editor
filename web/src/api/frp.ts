import { request } from "../utils/http";

export const getForwards = (params: any) => {
  return request.get("/frp/list", { params });
};

export const delForward = (data: any) => {
  return request.post("/frp/delete", data);
};

export const updateForward = (data: any) => {
  return request.post("/frp/update", data);
};

export const addForward = (data: any) => {
  return request.post("/frp/add", data);
};
