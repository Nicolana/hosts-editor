<script lang="ts" setup>
import { TabPanelName } from "element-plus";
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";

const router = useRouter();
const routes = useRoute();

// console.log("routes =", );
const current = ref<TabPanelName>(routes.path.valueOf());

const onTabChange = (name: TabPanelName) => {
  router.push(name as string);
  current.value = name;
};
</script>

<template>
  <el-tabs
    v-model="current"
    type="border-card"
    class="demo-tabs"
    @tab-change="onTabChange"
  >
    <el-tab-pane name="/">
      <template #label>
        <span class="custom-tabs-label">
          <el-icon>
            <calendar />
          </el-icon>
          <span>Hosts管理</span>
        </span>
      </template>
    </el-tab-pane>
    <el-tab-pane name="/frp">
      <template #label>
        <span class="custom-tabs-label">
          <el-icon>
            <DataAnalysis />
          </el-icon>
          <span>Frp端口映射</span>
        </span>
      </template>
    </el-tab-pane>
    <router-view />
  </el-tabs>
</template>

<style lang="scss" scoped>
.custom-tabs-label {
  display: flex;
  align-items: center;
  justify-content: space-between;

  i {
    margin-right: 5px;
  }
}
</style>
