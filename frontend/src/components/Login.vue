<script setup>
import {ref} from 'vue'
import ax from '../pkg/axios'
import Cookies from "js-cookie";
import {Nerr, Nsucc} from '../pkg/notify'

let form = ref({})

async function login() {
  if (!form.value.username) {
    Nerr("用户名不能为空")
    return
  }
  if (!form.value.password) {
    Nerr("密码不能为空")
    return
  }
  let da = await ax({
    url: "system/login",
    method: "post",
    data: form.value
  })
  if (!da) {
    return
  }
  Cookies.set("token", da.token)
  Nsucc("welcom " + da.nickname)
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
