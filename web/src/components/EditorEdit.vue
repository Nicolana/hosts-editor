<template>
  <el-dialog v-model="visible" title="Hosts编辑" append-to-body width="500px">
    <el-form :model="form" label-width="auto" label-position="top">
      <el-form-item label="域名" prop="hosts" required>
        <el-input v-model="form.hosts" autocomplete="off" />
      </el-form-item>
      <el-form-item label="IP" prop="ip" required>
        <el-input v-model="form.ip" autocomplete="off" />
      </el-form-item>
      <el-form-item label="备注" prop="comments">
        <el-input v-model="form.comments" autocomplete="off" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="onCancel">
          取消
        </el-button>
        <el-button type="primary" @click="onConfirm" :loading="submitting">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ElMessage } from "element-plus";
import { reactive, ref, watchEffect } from "vue";
import { StatusCode } from "../utils/consts";
import { request } from "../utils/http";

const visible = ref(false);
const submitting = ref(false);

const props = defineProps({
  rowInfo: {
    type: Object,
    required: true
  }
})

const form = reactive({
  hosts: '',
  ip: '',
  comments: '',
})

const clearForm = () => {
  form.hosts = '';
  form.ip = '';
  form.comments = '';
}

const toggleVisible = () => {
  visible.value = !visible.value
}

watchEffect(() => {
  form.hosts = props.rowInfo.hosts;
  form.ip = props.rowInfo.ip;
  form.comments = props.rowInfo.comments;
})

const onCancel = () => {
  visible.value = false;
  emit('onCancel')
}

const onConfirm = async () => {
  const payload = {
    row: props.rowInfo.index,
    hosts: form.hosts,
    ip: form.ip,
    comments: form.comments || undefined
  }
  submitting.value = true;
  try {
    const { data: res } = await request.post("/api/update", payload);
    console.log("update res = ", res)
    if (res.code === StatusCode.Success) {
      ElMessage.success(res.message)
    } else {
      ElMessage.error(res.message)
    }
  } catch (error) {
    console.error(error);
  }
  submitting.value = false;
  visible.value = false;
  emit('onConfirm')
}

const emit = defineEmits(['onCancel', 'onConfirm'])

defineExpose({
  visible,
  toggleVisible,
})
</script>

<style scoped>
</style>