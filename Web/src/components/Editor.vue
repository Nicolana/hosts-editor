<script setup lang="ts">
import { ref } from 'vue';

const tableData = [
    {
        date: '2016-05-03',
        name: 'Tom',
        state: 'California',
        city: 'Los Angeles',
        address: 'No. 189, Grove St, Los Angeles',
        zip: 'CA 90036',
        tag: 'Home',
    },
    {
        date: '2016-05-02',
        name: 'Tom',
        state: 'California',
        city: 'Los Angeles',
        address: 'No. 189, Grove St, Los Angeles',
        zip: 'CA 90036',
        tag: 'Office',
    },
    {
        date: '2016-05-04',
        name: 'Tom',
        state: 'California',
        city: 'Los Angeles',
        address: 'No. 189, Grove St, Los Angeles',
        zip: 'CA 90036',
        tag: 'Home',
    },
    {
        date: '2016-05-01',
        name: 'Tom',
        state: 'California',
        city: 'Los Angeles',
        address: 'No. 189, Grove St, Los Angeles',
        zip: 'CA 90036',
        tag: 'Office',
    },
]

const handleClick = () => { }
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
            <el-table-column prop="date" label="Date" />
            <el-table-column prop="name" label="Name" />
            <el-table-column prop="state" label="State" />
            <el-table-column prop="city" label="City" />
            <el-table-column prop="address" label="Address" />
            <el-table-column prop="zip" label="Zip" />
            <el-table-column label="Operations">
                <template #default>
                    <el-button link type="primary" size="small" @click="handleClick">Detail</el-button>
                    <el-button link type="primary" size="small">Edit</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-card>
</template>

<style scoped>
</style>
