<script setup lang="ts">
import { ElMessage } from 'element-plus';
import { ref } from 'vue';
import { StatusCode } from '../utils/consts';
import { request } from '../utils/http';
import EditorEdit from './EditorEdit.vue';

const tableData = ref([])
const rowInfo = ref({})
const modalRef = ref(null)

const loadList = () => {
  request.get("/api/list").then(({ data: res }) => {
    tableData.value = res.data;
  })
}

loadList();

const onEdit = (row: any) => {
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
</script>

<template>
  <el-card class="box-card" :body-style="{ padding: 0 }">
    <template #header>
      <div class="card-header">
        <span>Hosts管理</span>
        <div></div>
      </div>
    </template>
    <el-table :data="tableData" style="width: 100%" size="large">
      <el-table-column prop="index" label="行数" />
      <el-table-column prop="hosts" label="域名" />
      <el-table-column prop="ip" label="IP" />
      <el-table-column label="Operations">
        <template #default="{ row }">
          <el-button type="primary" @click="onEdit(row)">修改</el-button>
          <el-popconfirm title="确定删除该DNS解析?" @confirm="onDelete(row.index)">
            <template #reference>
              <el-button type="danger">删除</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
  <EditorEdit ref="modalRef" :rowInfo="rowInfo" @onCancel="onCancel" />
</template>

<style scoped>
</style>
