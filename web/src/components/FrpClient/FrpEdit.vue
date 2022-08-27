<template>
  <el-dialog v-model="visible" title="FRP配置添加" append-to-body width="800px">
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
            <!-- Switch类型 -->
            <el-form-item
              :label="row.name"
              :prop="`frpConfigs[${index}].value`"
              :rules="[{ required: true, message: `请选择${row.key}` }]"
              layout="inline"
              v-else-if="row.type === FrpFormItemType.switch"
            >
              <el-switch v-model="row.value" />
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
            <el-dropdown>
              <el-button type="primary">
                添加预设配置<el-icon class="el-icon--right"
                  ><arrow-down
                /></el-icon>
              </el-button>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item
                    v-for="(item, index) in presetConfig.arr"
                    :key="index"
                    @click="onPresetAdd(index)"
                    >{{ item.name }}</el-dropdown-item
                  >
                </el-dropdown-menu>
              </template>
            </el-dropdown>
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
import { FormInstance, messageConfig } from "element-plus";
import { ElMessage } from "element-plus";
import { computed, reactive, ref, watchEffect } from "vue";
import { ModalStatusCode, StatusCode } from "@/utils/consts";
import { addForward, updateForward } from "@/api/frp";
import { CirclePlus, Delete } from "@element-plus/icons-vue";
import {
  FrpConfigItemType,
  FrpFormItemType,
  GetFormItemByKey,
  FrpClientConfig,
} from "./consts";

const visible = ref(false);
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

// 预设配置
const presetConfig = reactive({
  arr: [
    {
      name: "普通TCP/UDP服务",
      keys: ["type", "local_ip", "local_port", "remote_port"],
    },
    {
      name: "WEB服务",
      keys: [
        "type",
        "local_ip",
        "local_port",
        "remote_port",
        "use_encryption",
        "use_compression",
        "http_user",
        "http_pwd",
        "subdomain",
        "custom_domains",
        "locations",
      ],
    },
    {
      name: "P2P-TCP",
      keys: [
        "type",
        "sk",
        "local_ip",
        "local_port",
        "use_encryption",
        "use_compression",
      ],
    },
  ],
});

const onPresetAdd = (index: number) => {
  form.frpConfigs = presetConfig.arr[index].keys.map((key) =>
    GetFormItemByKey(key)
  );
};

const addRow = () => {
  const item = FrpClientConfig.find((conf) => conf.key === addType.value);
  if (item) {
    form.frpConfigs.push({ ...item });
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
  if (props.rowInfo.configs) {
    form.frpConfigs = props.rowInfo.configs;
  } else {
    form.frpConfigs = [GetFormItemByKey("type")];
  }
  form.name = props.rowInfo.name;
});

const onCancel = () => {
  visible.value = false;
  emit("onCancel");
};

const onConfirm = async () => {
  submitting.value = true;
  const payload: { [keyProp: string]: any } = {
    name: form.name,
  };

  for (let conf of form.frpConfigs) {
    payload[conf.key] = conf.value;
  }
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
