<script setup lang="ts">
import { ElMessage, messageConfig } from 'element-plus';
import { ref } from 'vue';
import { getFrpStatus, restartFrp, startFrp, stopFrp } from '../../api/frp';
import { FrpExecStatus, StatusCode } from '../../utils/consts';

const status = ref(FrpExecStatus.Stopped)

const getStatus = async () => {
    const { data: res } = await getFrpStatus();
    if (res.code === StatusCode.Success) {
        status.value = +res.data?.status;
    } else {
        ElMessage.error(res.message)
    }
}

const start = async () => {
    const { data: res } = await startFrp();
    if (res.code === StatusCode.Success) {
        status.value = FrpExecStatus.Running;
        ElMessage.success(res.message)
        await getStatus()
    } else {
        ElMessage.error(res.message)
    }
}

const stop = async () => {
    const { data: res } = await stopFrp();
    if (res.code === StatusCode.Success) {
        status.value = FrpExecStatus.Stopped;
        ElMessage.success(res.message)
        await getStatus()
    } else {
        ElMessage.error(res.message)
    }
}

const restart = async () => {
    const { data: res } = await restartFrp();
    if (res.code === StatusCode.Success) {
        status.value = FrpExecStatus.Running;
        ElMessage.success(res.message)
        await getStatus()
    } else {
        ElMessage.error(res.message)
    }
}

getStatus();

setInterval(() => {
    getStatus();
}, 1000 * 2)
</script>

<template>
    <el-card class="box-card" title="FRP运行状态">
        <template #header>
            <div class="card-header">
                <span>FRP客户端运行状态</span>
            </div>
        </template>
        <div class="frp-status-wrapper">
            <div class="left">
                <el-space>
                    <el-tag class="ml-2" type="info" v-if="status === FrpExecStatus.Stopped">未启动</el-tag>
                    <el-tag class="ml-2" type="success" v-else>运行中</el-tag>
                </el-space>
            </div>
            <div class="right">
                <el-button type="primary" @click="restart">
                    重启
                </el-button>
                <el-button type="danger" v-if="status === FrpExecStatus.Running" @click="stop">
                    停止
                </el-button>
                <el-button type="success" v-else @click="start">
                    启动
                </el-button>
            </div>
        </div>
    </el-card>
</template>

<style lang="scss" scoped>
.box-card {
    margin-bottom: 20px;

    .card-header {
        height: 20px;
        display: flex;
        align-items: center;
        font-size: 12px;
        color: var(--el-color-info);
    }
}

.frp-status-wrapper {
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: space-between;
}
</style>
