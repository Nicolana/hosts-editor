<template>
    <el-dialog v-model="visible" title="Hosts编辑" append-to-body width="500px">
        <el-form ref="formRef" :model="form" label-width="auto" label-position="top">
            <el-form-item label="服务器地址" prop="server_addr" :rules="[{ required: true, message: '请输入服务器地址' }]">
                <el-input v-model="form.server_addr" autocomplete="off" />
            </el-form-item>
            <el-form-item label="服务器端口" prop="server_port" :rules="[{ required: true, message: '请输入服务器端口' }]">
                <el-input v-model="form.server_port" autocomplete="off" />
            </el-form-item>
            <el-form-item label="Token" prop="token" :rules="[{ required: true, message: '请输入校验token' }]">
                <el-input v-model="form.token" autocomplete="off" />
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
import { StatusCode } from "@/utils/consts";
import { FrpServerTypes } from "@/utils/types";
import type { FormInstance } from 'element-plus'
import { updateServerConfig } from "@/api/frp";

const visible = ref(false);
const submitting = ref(false);
const formRef = ref<FormInstance>();

const props = defineProps({
    rowInfo: {
        type: Object,
        required: true,
    },
});

const form = reactive({
    server_addr: '',
    server_port: '', // 0 ~ 65535
    token: '',
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
    form.server_addr = props.rowInfo.server_addr;
    form.server_port = props.rowInfo.server_port;
    form.token = props.rowInfo?.token;
});

const onCancel = () => {
    visible.value = false;
    emit("onCancel");
};

const onConfirm = async () => {
    const payload: FrpServerTypes = {
        server_addr: form.server_addr,
        server_port: +form.server_port,
        token: form.token,
    };
    submitting.value = true;
    try {
        const { data: res } = await updateServerConfig(payload);
        if (res.code === StatusCode.Success) {
            ElMessage.success(res.message);
        } else {
            ElMessage.error(res.message);
        }
    } catch (error) {
        console.error(error);
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
