<template>
  <el-dialog v-model="visible" title="FRP配置添加" append-to-body>
    <el-form
      ref="formRef"
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
          style="width: 100%"
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
      <div v-for="(row, index) in form.frpConfigs" :key="index">
        <el-row :gutter="10">
          <el-col :span="11">
            <el-form-item
              label="配置类型"
              prop="name"
              :rules="[{ required: true, message: '请选择配置类型' }]"
            >
              <el-select
                v-model="row.name"
                style="width: 100%"
                placeholder="请选择配置类型"
              >
                <el-option
                  v-for="tt in frpConfigTypes"
                  :key="tt.value"
                  :label="tt.label"
                  :value="tt.value"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item
              label="配置值"
              prop="value"
              :rules="[{ required: true, message: '请输入配置值' }]"
            >
              <el-input v-model="row.value" autocomplete="off" />
            </el-form-item>
          </el-col>
          <el-col :span="2">
            <el-form-item label="操作">
              <el-button
                type="danger"
                :icon="Delete"
                circle
                @click="deleteRow(index)"
              />
            </el-form-item>
          </el-col>
        </el-row>
      </div>
      <el-form-item>
        <el-row :gutter="5">
          <el-col :span="10">
            <el-select
              v-model="addType"
              filterable
              default-first-option
              placeholder="请选择配置参数"
            >
              <el-option
                v-for="item in options"
                filterable
                :key="item.name"
                :label="item.name"
                :value="item.name"
              />
            </el-select>
          </el-col>
          <el-col :span="12">
            <el-button :icon="CirclePlus" @click="addRow">
              添加新的配置项
            </el-button>
          </el-col>
        </el-row>
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
import type { FormInstance } from "element-plus";
import { ElMessage } from "element-plus";
import { computed, reactive, ref, watchEffect } from "vue";
import {
  DEFAULT_IP_ADDRESS,
  ModalStatusCode,
  NetworkType,
  StatusCode,
} from "@/utils/consts";
import { FrpPayloadTypes } from "@/utils/types";
import { addForward, updateForward } from "@/api/frp";
import { CirclePlus, Delete } from "@element-plus/icons-vue";
import {
  FrpConfigItemType,
  FrpFormItemEnum,
  FrpFormItemType,
  frpTypes,
  FrpClientConfig,
} from "./consts";

const visible = ref(true);
const submitting = ref(false);
const formRef = ref<FormInstance>();
const addType = ref("");

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
  // type: NetworkType.TCP,
  // local_ip: DEFAULT_IP_ADDRESS,
  // local_port: "", // 0 ~ 65535
  // remote_port: "",
  // plugin: "",
  // sk: "",
  // plugin_user: "",
  // plugin_passwd: "",
  // bandwidth_limit: "",
  // use_encryption: false,
  // use_compression: false,
  frpConfigs: [
    {
      name: "name",
      value: "",
      type: FrpFormItemType.text,
    },
  ] as FrpConfigItemType[],
});

const options = computed(() =>
  FrpClientConfig.filter(
    (itme) => form.frpConfigs.findIndex((row) => row.name !== itme.name) > -1
  )
);

const addRow = () => {
  form.frpConfigs.push({ name: "", value: "", type: FrpFormItemType.text });
};

const deleteRow = (index: number) => {
  form.frpConfigs.splice(index, 1);
};

const frpConfigTypes = ref([
  {
    label: "plugin",
    value: "plugin",
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
