<template>
  <el-dialog v-model="visible" title="Hosts编辑" append-to-body>
    <el-form
      ref="formRef"
      :inline="true"
      :model="form"
      label-width="auto"
      label-position="top"
    >
      <el-form-item
        label="名称"
        prop="name"
        :rules="[{ required: true, message: '请输入名称' }]"
      >
        <el-input v-model="form.name" autocomplete="off" />
      </el-form-item>
      <el-form-item
        label="类型"
        prop="type"
        :rules="[{ required: true, message: '请选择类型' }]"
      >
        <el-select
          v-model="form.type"
          placeholder="请选择类型"
          default-first-option
        >
          <el-option
            v-for="tp in frpTypes"
            :key="tp.value"
            :label="tp.label"
            :value="tp.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="本地IP地址"
        prop="local_ip"
        :rules="[{ required: true, message: '请输入需要映射的本地IP地址' }]"
      >
        <el-input v-model="form.local_ip" autocomplete="off" />
      </el-form-item>
      <el-form-item
        label="本地端口"
        prop="local_port"
        :rules="[{ required: true, message: '请输入需要映射的本地端口' }]"
      >
        <el-input v-model="form.local_port" autocomplete="off" />
      </el-form-item>
      <el-form-item label="服务器端口" prop="remote_port">
        <el-input v-model="form.remote_port" autocomplete="off" />
      </el-form-item>
      <el-form-item label="加密传输" prop="use_encryption">
        <el-radio-group v-model="form.use_encryption" class="ml-4">
          <el-radio label="1" size="large">Option 1</el-radio>
          <el-radio label="2" size="large">Option 2</el-radio>
        </el-radio-group>
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
import {
  DEFAULT_IP_ADDRESS,
  ModalStatusCode,
  NetworkType,
  StatusCode,
} from "../../utils/consts";
import { FrpPayloadTypes, PayloadTypes } from "../../utils/types";
import type { FormInstance } from "element-plus";
import { addForward, updateForward } from "../../api/frp";

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
  name: "",
  type: NetworkType.TCP,
  local_ip: DEFAULT_IP_ADDRESS,
  local_port: "", // 0 ~ 65535
  remote_port: "",
  use_encryption: false,
});

const frpTypes = ref([
  {
    label: "tcp",
    value: "tcp",
  },
  {
    label: "udp",
    value: "udp",
  },
  {
    label: "xtcp",
    value: "xtcp",
  },
  {
    label: "tcpmux",
    value: "tcpmux",
  },
  {
    label: "stcp",
    value: "stcp",
  },
  {
    label: "https",
    value: "https",
  },
  {
    label: "http",
    value: "http",
  },
]);

const clearForm = (el: FormInstance | undefined) => {
  if (!el) return;
  el.resetFields();
  el.clearValidate();
};

const toggleVisible = () => {
  visible.value = !visible.value;
};

watchEffect(() => {
  form.name = props.rowInfo.name;
  form.type = props.rowInfo.type;
  form.local_ip = props.rowInfo?.local_ip || DEFAULT_IP_ADDRESS;
  form.local_port = props.rowInfo?.local_port ?? "";
  form.remote_port = props.rowInfo?.remote_port ?? "";
});

const onCancel = () => {
  visible.value = false;
  emit("onCancel");
};

const onConfirm = async () => {
  const payload: FrpPayloadTypes = {
    name: form.name,
    type: form.type,
    local_ip: form.local_ip,
    local_port: +form.local_port,
    remote_port: +form.remote_port,
  };
  submitting.value = true;
  if (props.modalStatus === ModalStatusCode.Edit) {
    try {
      const { data: res } = await updateForward(payload);
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
      const { data: res } = await addForward(payload);
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

<style scoped></style>
