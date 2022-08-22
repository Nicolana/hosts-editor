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

const menus = ref([
  {
    path: "/",
    name: "Hosts管理",
    icon: "Calendar",
  },
  {
    path: "/frp",
    name: "Frp端口映射",
    icon: "DataAnalysis",
  },
  {
    path: "/base64",
    name: "Base64加解码",
    icon: "DataAnalysis",
  },
]);
</script>

<template>
  <el-tabs
    v-model="current"
    type="border-card"
    class="demo-tabs"
    @tab-change="onTabChange"
  >
    <el-tab-pane v-for="item in menus" :key="item.path" :name="item.path">
      <template #label>
        <span class="custom-tabs-label">
          <el-icon>
            <component :is="item.icon" />
          </el-icon>
          <span>{{ item.name }}</span>
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
