<template>
  <el-dialog v-model="visible" title="Hosts编辑" append-to-body width="500px">
    <el-form ref="formRef" :model="form" label-width="auto" label-position="top">
      <el-form-item label="域名" prop="hosts" :rules="[{ required: true, message: '请输入要解析的域名' }]">
        <el-input v-model="form.hosts" autocomplete="off" />
      </el-form-item>
      <el-form-item label="IP" prop="ip" :rules="[{ required: true, message: '请输入IP地址' }]">
        <el-input v-model="form.ip" autocomplete="off" />
      </el-form-item>
      <el-form-item label="备注" prop="comments">
        <el-input v-model="form.comments" autocomplete="off" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="onCancel"> 取消 </el-button>
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
import { DEFAULT_IP_ADDRESS, ModalStatusCode, StatusCode } from "../../utils/consts";
import { PayloadTypes } from "../../utils/types";
import type { FormInstance } from 'element-plus'
import { addHosts, updateHosts } from "../../api/hosts";

const visible = ref(false);
const submitting = ref(false);
const formRef = ref<FormInstance>();

const props = defineProps({
  rowInfo: {
    type: Object,
    required: true,
  },
  modalStatus: {
    type: Number,
    default: ModalStatusCode.Create,
  },
});

const form = reactive({
  hosts: "",
  ip: DEFAULT_IP_ADDRESS,
  comments: "",
});

const clearForm = (el: FormInstance | undefined) => {
  if (!el) return;
  el.resetFields();
  el.clearValidate();
};

const toggleVisible = () => {
  visible.value = !visible.value;
};

watchEffect(() => {
  form.hosts = props.rowInfo.hosts;
  form.ip = props.rowInfo?.ip || DEFAULT_IP_ADDRESS;
  form.comments = props.rowInfo.comments;
});

const onCancel = () => {
  visible.value = false;
  emit("onCancel");
};

const onConfirm = async () => {
  const payload: PayloadTypes = {
    hosts: form.hosts,
    ip: form.ip,
    comments: form.comments || undefined,
  };
  submitting.value = true;
  if (props.modalStatus === ModalStatusCode.Edit) {
    payload.row = props.rowInfo.index;
    try {
      const { data: res } = await updateHosts(payload);
      if (res.code === StatusCode.Success) {
        ElMessage.success(res.message);
      } else {
        ElMessage.error(res.message);
      }
    } catch (error) {
      console.error(error);
    }
  } else {
    try {
      const { data: res } = await addHosts(payload);
      if (res.code === StatusCode.Success) {
        ElMessage.success(res.message);
      } else {
        ElMessage.error(res.message);
      }
    } catch (error) {
      console.error(error);
    }
  }
  submitting.value = false;
  visible.value = false;
  clearForm(formRef.value);
  emit("onConfirm");
};

const emit = defineEmits(["onCancel", "onConfirm"]);

defineExpose({
  visible,
  toggleVisible,
});
</script>

<style scoped>
</style>
