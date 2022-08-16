<script setup>
import {RouterView} from 'vue-router'
import Top from './Top.vue'
import Left from './Left.vue'</script>

<template class="app">
    <el-container class="container back">
        <el-header class="header back">
            <div class="top">
                <Top></Top>
            </div>
        </el-header>
        <el-container>
            <el-aside width="200px" class="aside back">
                <Left class="left back"></Left>
            </el-aside>
            <el-main class="main">
                <router-view></router-view>
            </el-main>
        </el-container>
    </el-container>
</template>

<style>
.app {
    width: 100%;
    height: 100%;
}

.container {
    height: 100%;
}

.back {
    background: #494986;
}

.aside {
    padding-left: 10px;
    height: 100%;
    width: 160px;
}

.left {
    height: 100%;
}

.top {
    float: right;
}

.header {
    width: 100%;
}

.el-submenu__title:hover, .el-menu-item:focus, .el-menu-item:hover {
    background-color: #494986;
}

.el-menu {
    border-right: none;
}

.main {
    background: #b4becb;
    border-radius: 10px;
}

</style>
