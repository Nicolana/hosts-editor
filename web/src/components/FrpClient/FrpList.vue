<script setup lang="ts">
import { ElMessage } from "element-plus";
import { Search } from "@element-plus/icons-vue";
import { computed, reactive, ref } from "vue";
import { ModalStatusCode, StatusCode } from "../../utils/consts";
import EditorEdit from "./FrpEdit.vue";
import { FrpServerTypes, PaginationType } from "../../utils/types";
import { getForwards, delForward, getServerConfig } from "../../api/frp";
import FrpStatus from "./FrpStatus.vue";
import FrpServerEdit from "./FrpServerEdit.vue";
import useFrpStore from "@/store/useFrpStore";
import { FrpConfigItemType, FrpFormItemType, GetFormItemByKey } from "./consts";

const tableData = ref([]);
const rowInfo = ref<{ [key: string]: any }>({});
const modalRef = ref(null);
const serverModalRef = ref();
const modalStatus = ref(ModalStatusCode.Create); // 默认是创建
const searchText = ref("");
const pagination = reactive<PaginationType>({
  page: 1,
  size: 20,
  total: 0,
});
const frpStore = useFrpStore();
const serverConfig = computed(() => frpStore.server);

const loadServerConfig = async () => {
  const { data: res } = await getServerConfig();
  if (res.code === StatusCode.Success) {
    frpStore.updateServer(res.data);
  }
};
loadServerConfig();

const loadList = () => {
  const params: Partial<PaginationType> = {
    page: pagination.page,
    size: pagination.size,
  };
  if (searchText.value) {
    params.search = searchText.value;
  }

  getForwards(params).then(({ data: res }) => {
    if (res.code === StatusCode.Success) {
      const { total, data: list } = res.data;
      pagination.total = total;
      tableData.value = list;
    }
  });
};

loadList();

const onCreate = () => {
  modalStatus.value = ModalStatusCode.Create;
  // @ts-ignore
  modalRef.value?.toggleVisible();
};

const onEdit = (row: any) => {
  modalStatus.value = ModalStatusCode.Edit;
  const payload: { [key: string]: any } = { ...rowInfo };
  payload.name = row.name;
  const frpConfigs: FrpConfigItemType[] = [];
  for (const key of Object.keys(row)) {
    if (key === "name") {
      continue;
    }
    const item = GetFormItemByKey(key);
    if (item) {
      if (item.type === FrpFormItemType.number) {
        item.value = +row[key];
      } else {
        item.value = row[key];
      }
      frpConfigs.push(item);
    }
  }
  payload.configs = frpConfigs;
  rowInfo.value = payload;
  // @ts-ignore
  modalRef.value?.toggleVisible();
};

const onCancel = () => {
  rowInfo.value = {};
};

const onDelete = async (name: string) => {
  try {
    const { data: res } = await delForward({ name });
    if (res.code === StatusCode.Success) {
      ElMessage.success(res.message);
      loadList();
    }
  } catch (error) {
    console.error(error);
  }
};

const onSearchChange = (val: string) => {
  searchText.value = val;
  loadList();
};

const onServerEdit = () => {
  serverModalRef.value.toggleVisible();
};
</script>

<template>
  <FrpStatus />
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <div class="search-wrapper">
          <el-input
            v-model="searchText"
            class="w-50 m-2"
            placeholder="请输入名称搜索"
            :prefix-icon="Search"
            @input="onSearchChange"
          />
        </div>
        <div>
          <el-button @click="onServerEdit">服务器配置</el-button>
          <el-button type="primary" @click="onCreate">新增</el-button>
        </div>
      </div>
    </template>
    <el-table :data="tableData" style="width: 100%" size="large">
      <el-table-column prop="name" width="140" label="名称" />
      <el-table-column prop="type" label="类型" />
      <el-table-column prop="local_ip" label="本地IP" />
      <el-table-column prop="local_port" label="本地端口" />
      <el-table-column prop="remote_port" label="服务器端口" />
      <el-table-column label="操作" width="200">
        <template #default="{ row }">
          <el-button type="primary" @click="onEdit(row)">修改</el-button>
          <el-popconfirm
            title="确定删除该端口映射么?"
            confirm-button-text="确定"
            cancel-button-text="取消"
            @confirm="onDelete(row.name)"
          >
            <template #reference>
              <el-button type="danger">删除</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination-wrapper">
      <el-pagination
        background
        layout="prev, pager, next"
        hide-on-single-page
        :total="pagination.total"
        :page-size="pagination.size"
      />
    </div>
  </el-card>
  <EditorEdit
    ref="modalRef"
    :rowInfo="rowInfo"
    @on-cancel="onCancel"
    @on-confirm="loadList"
    :modalStatus="modalStatus"
  />
  <FrpServerEdit
    ref="serverModalRef"
    :row-info="serverConfig"
    @on-confirm="loadServerConfig"
  />
</template>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.title-search {
  display: flex;
  align-items: center;
}

.search-wrapper {
  margin-left: 10px;
}

.pagination-wrapper {
  padding: 20px 0;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
