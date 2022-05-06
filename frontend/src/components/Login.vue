<script setup>
import {ref} from 'vue'
import ax, {handleError} from '../pkg/axios'

let form = ref({})

function login() {
  ax.post('system/login', form.value)
      .then(function (response) {
        if (response.data.errcode === 0) {
          alert("succ")
        } else {
          alert(response.data.errmsg)
        }
      })
      .catch(handleError);
}

</script>

<template>

  <el-form :model="form" label-width="80px">
    <el-form-item label="用户名">
      <el-input v-model="form.username"></el-input>
    </el-form-item>
    <el-form-item label="密码">
      <el-input v-model="form.password" type="password"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="login">登陆</el-button>
    </el-form-item>
  </el-form>

</template>

<style scoped>

</style>
