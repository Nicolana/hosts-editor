<script setup lang="ts">
import { ElMessage } from 'element-plus';
import { Search } from '@element-plus/icons-vue'
import { reactive, ref } from 'vue';
import { ModalStatusCode, StatusCode } from '../utils/consts';
import { request } from '../utils/http';
import EditorEdit from './EditorEdit.vue';
import { PaginationType } from '../utils/types';

const tableData = ref([])
const rowInfo = ref({})
const modalRef = ref(null)
const modalStatus = ref(ModalStatusCode.Create); // 默认是创建
const searchText = ref('');
const pagination = reactive<PaginationType>({
  page: 1,
  size: 20,
  total: 0,
})

const loadList = () => {
  const params: Partial<PaginationType> = {
    page: pagination.page,
    size: pagination.size
  }
  if (searchText.value) {
    params.search = searchText.value
  }

  request.get("/api/list", { params }).then(({ data: res }) => {
    if (res.code === StatusCode.Success) {
      const { total, data: list } = res.data;
      pagination.total = total;
      tableData.value = list;
    }
  })
}

loadList();

const onCreate = () => {
  modalStatus.value = ModalStatusCode.Create;
  // @ts-ignore
  modalRef.value?.toggleVisible();
}

const onEdit = (row: any) => {
  modalStatus.value = ModalStatusCode.Edit;
  rowInfo.value = row;
  // @ts-ignore
  modalRef.value?.toggleVisible();
}

const onCancel = () => {
  rowInfo.value = {}
}

const onDelete = async (rowId: number) => {
  try {
    const { data: res } = await request.post("/api/delete", { row: rowId })
    if (res.code === StatusCode.Success) {
      ElMessage.success(res.message);
      loadList();
    }
  } catch (error) {
    console.error(error)
  }
}

const onSearchChange = (val: string) => {
  searchText.value = val;
  loadList();
}
</script>

<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <div class="title-search">
          <span>Hosts管理</span>
        </div>
        <div class="search-wrapper">
          <el-input v-model="searchText" class="w-50 m-2" size="small" placeholder="请输入域名" :prefix-icon="Search"
            @input="onSearchChange" />
        </div>
        <el-button type="primary" @click="onCreate">新增</el-button>
      </div>
    </template>
    <el-table :data="tableData" style="width: 100%" size="large">
      <el-table-column prop="index" width="100" label="行数" />
      <el-table-column prop="hosts" label="域名" />
      <el-table-column prop="ip" label="IP" />
      <el-table-column label="操作" width="200">
        <template #default="{ row }">
          <el-button type="primary" @click="onEdit(row)">修改</el-button>
          <el-popconfirm title="确定删除该DNS解析?" confirm-button-text="确定" cancel-button-text="取消"
            @confirm="onDelete(row.index)">
            <template #reference>
              <el-button type="danger">删除</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination-wrapper">
      <el-pagination background layout="prev, pager, next" hide-on-single-page :total="pagination.total"
        :page-size="pagination.size" />
    </div>
  </el-card>
  <EditorEdit ref="modalRef" :rowInfo="rowInfo" @on-cancel="onCancel" @on-confirm="loadList"
    :modalStatus="modalStatus" />
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
