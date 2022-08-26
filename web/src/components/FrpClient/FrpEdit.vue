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
      <div v-for="(row, index) in form.frpConfigs" :key="index">
        <el-row :gutter="10">
          <el-col :span="21">
            <!-- 文本类型 -->
            <el-form-item
              :label="row.name"
              :prop="`frpConfigs[${index}].value`"
              :rules="[{ required: true, message: `请输入${row.key}` }]"
              v-if="row.type === FrpFormItemType.text"
            >
              <el-input v-model="row.value" autocomplete="off" />
            </el-form-item>
            <!-- 数值类型 -->
            <el-form-item
              :label="row.name"
              :prop="`frpConfigs[${index}].value`"
              :rules="[{ required: true, message: `请输入${row.key}` }]"
              v-if="row.type === FrpFormItemType.number"
            >
              <el-input-number
                v-model="row.value"
                type="number"
                :min="row.min"
                :max="row.max"
                autocomplete="off"
              />
            </el-form-item>
            <!-- 枚举类型 -->
            <el-form-item
              :label="row.name"
              :prop="`frpConfigs[${index}].value`"
              :rules="[{ required: true, message: `请选择${row.key}` }]"
              v-else-if="row.type === FrpFormItemType.enum"
            >
              <el-select
                v-model="row.value"
                :placeholder="`请选择${row.key}`"
                style="width: 100%"
                default-first-option
              >
                <el-option
                  v-for="(enumValue, enumKey) in row.valueEnum"
                  :key="enumKey"
                  :label="enumValue.text"
                  :value="enumKey"
                />
              </el-select>
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
      <!-- RecordCreator -->
      <el-form-item>
        <el-row :gutter="5">
          <el-col :span="8">
            <el-select
              v-model="addType"
              filterable
              default-first-option
              placeholder="请选择配置参数"
            >
              <el-option
                v-for="item in options"
                filterable
                :key="item.key"
                :label="item.name"
                :value="item.key"
              />
            </el-select>
          </el-col>
          <el-col :span="8">
            <el-button :icon="CirclePlus" @click="addRow">
              添加新的配置项
            </el-button>
          </el-col>
          <el-col :span="8">
            <el-button type="primary" @click="addRow">
              一键添加预设配置
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
  StatusCode,
} from "@/utils/consts";
import { FrpPayloadTypes } from "@/utils/types";
import { addForward, updateForward } from "@/api/frp";
import { CirclePlus, Delete } from "@element-plus/icons-vue";
import {
  FrpConfigItemType,
  FrpFormItemType,
  NetworkType,
  GetFormItemByKey,
  FrpClientConfig,
} from "./consts";
import { debug } from "console";

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
  frpConfigs: [
    GetFormItemByKey("type"), // 设置默认填类型
  ] as FrpConfigItemType[],
});

// 所有可选择的类型选项
const options = computed(() => {
  const data = FrpClientConfig.filter(
    (item) => form.frpConfigs.findIndex((row) => row.key === item.key) < 0
  );
  return data;
});

const addRow = () => {
  const item = FrpClientConfig.find((conf) => conf.key === addType.value);
  if (item) {
    form.frpConfigs.push(item);
    addType.value = "";
  }
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
