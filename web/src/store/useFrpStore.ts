import { FrpServerTypes } from "@/utils/types";
import { defineStore } from "pinia";

const useFrpStore = defineStore("frpStore", {
  state: () => {
    return {
      server: {} as FrpServerTypes,
    };
  },
  getters: {
    serverAddr: (state) => state.server.server_addr,
  },
  actions: {
    updateServer(server: FrpServerTypes) {
      this.server = server;
    },
  },
});

export default useFrpStore;
