<template>
  <div
    class="relative h-full w-full bg-no-repeat bg-cover bg-gray-100 overflow-y-auto wrap-login"
  >
    <base-wrapper class="flex-center-center">
      <div
        class="flex-c-center-center bg-color-white b-rd-10"
        style="height: 400px; width: 500px"
      >
        <h1 class="font-size-lg">Zone Admin</h1>
        <div class="logo">
          <!--<el-image :src="logoImage" class="logo-img" ></el-image>-->
        </div>
        <div class="m-t-20">
          <el-form ref="loginFormRef" :model="loginForm">
            <el-form-item
              prop="username"
              :rules="[
                { required: true, trigger: 'blur', message: '请输入账号' },
              ]"
            >
              <el-input
                size="large"
                v-model="loginForm.username"
                prefix-icon="User"
                placeholder="请输入账号"
                clearable
                maxlength="30"
                style="width: 300px"
              />
            </el-form-item>
            <el-form-item
              prop="password"
              :rules="[
                { required: true, trigger: 'blur', message: '请输入密码' },
              ]"
            >
              <el-input
                size="large"
                v-model="loginForm.password"
                prefix-icon="Lock"
                placeholder="请输入密码"
                show-password
                maxlength="50"
                style="width: 300px"
              />
            </el-form-item>
          </el-form>
          <div class="tips">
            <span>用户名: admin</span>
            <span class="m-l-20"> 密码: 123456</span>
          </div>
          <el-button
            size="large"
            type="primary"
            class="m-t-10 w-full"
            @click="handleLogin"
            >登 录
          </el-button>
        </div>
      </div>
      <div class="copyright">
        <p>FUCKING THE WORLD</p>
      </div>
    </base-wrapper>
  </div>
</template>

<script setup>
import { getCurrentInstance, proxyRefs } from "vue";

// 组件实例
const { proxy } = getCurrentInstance();
const { login } = proxy.$store.user.useUserStore();
const loginForm = $ref({});
const logoImage = new URL("../../assets/images/logo.png", import.meta.url).href;

/*
{
  code:200,
  msg:"ok",
  data: {
    tokeName: "auth-smallboot",
    tokenValue: "hhhhhhhhhh"
  }
}
 */
function handleLogin() {
  proxy.$refs.loginFormRef.validate((valid) => {
    if (valid) {
      login(loginForm).then(() => {
        let fullPath = proxy.$route.fullPath;
        if (fullPath.startsWith("/login?redirect=")) {
          let lastPath = fullPath.replace("/login?redirect=", "");
          // 跳转到上次退出的页面
          proxy.$router.push({ path: lastPath });
        } else {
          // 跳转到首页
          proxy.$router.push({ path: "/" });
        }
      });
    }
  });
}
</script>

<style lang="scss" scoped>
.wrap-login {
  background-image: url("@/assets/svg/login-bg.svg");
  background-repeat: no-repeat;
  background-position: center 110px;
  background-size: 100%;
  background-color: #f0f2f5;
}

.copyright {
  width: 100%;
  position: absolute;
  bottom: 20px;
  font-size: 14px;
  text-align: center;
  color: #ccc;
}
</style>
