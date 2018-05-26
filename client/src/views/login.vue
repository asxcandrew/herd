<template>
  <div id="login">
    <h1>{{ $t("views.login.title") }}</h1>
    <p>{{ $t("views.login.message") }}</p>
    <el-form ref="form" :model="form" :rules="rules">
      <el-form-item prop="email">
        <el-input v-model="form.email" :placeholder="$t('form.email.name')"></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input type="password" v-model="form.password" placeholder="Password" maxlength=20>
        </el-input>
      </el-form-item>
    </el-form>
    <p>
      {{ $t("views.login.signUpProposal") }}
      <el-button type="text" @click="loginDialogVisible = true">
        {{ $t("views.login.signUpButton") }}
      </el-button>.
    </p>
  </div>
</template>

<script>
export default {
  name: 'login',
  data() {
    return {
      form: {
        email: '',
        password: '',
      },
      rules: {
        email: [
          { required: true, message: this.$t('form.email.required'), trigger: 'blur' },
          { type: 'email', message: this.$t('form.email.incorrect'), trigger: 'blur' },
        ],
        password: [
          { required: true, message: this.$t('form.password.required'), trigger: 'blur' },
        ],
      },
    };
  },
};
</script>
