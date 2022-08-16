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

export const startFrp = () => {
  return request.get("/frp/start");
};

export const stopFrp = () => {
  return request.get("/frp/stop");
};

export const getFrpStatus = () => {
  return request.get("/frp/status");
};

export const restartFrp = () => {
  return request.get("/frp/restart");
};

export const getServerConfig = () => {
  return request.get("/frp/server");
};
export const updateServerConfig = (payload: any) => {
  return request.post("/frp/server", payload);
};
