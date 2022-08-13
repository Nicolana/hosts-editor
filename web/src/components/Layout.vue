<script lang="ts" setup>
import Editor from "./Editor.vue";

</script>

<template>
    <el-tabs type="border-card" class="demo-tabs">
        <el-tab-pane>
            <template #label>
                <span class="custom-tabs-label">
                    <el-icon>
                        <calendar />
                    </el-icon>
                    <span>Hosts管理</span>
                </span>
            </template>
            <Editor />
        </el-tab-pane>
        <el-tab-pane>
            <template #label>
                <span class="custom-tabs-label">
                    <el-icon>
                        <DataAnalysis />
                    </el-icon>
                    <span>Frp端口映射</span>
                </span>
            </template>
            FRP 端口映射工具开发
        </el-tab-pane>
    </el-tabs>
</template>

<style lang="scss" scoped>
.custom-tabs-label {
    display: flex;
    align-items: center;
    justify-content: space-between;

    i {
        margin-right: 5px;
    }
}
</style>