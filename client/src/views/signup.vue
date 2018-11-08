<template>
  <div id="signup">
    <h1>{{ $t("views.signup.title") }}</h1>
    <p>{{ $t("views.signup.message") }}</p>
    <el-form ref="form" :model="form" :rules="rules">
      <el-form-item prop="username">
        <el-input v-model="form.username" :placeholder="$t('form.username.name')">
          <template slot="append">@{{ domainName }}</template>
        </el-input>
      </el-form-item>
      <el-form-item prop="email">
        <el-input v-model="form.email" :placeholder="$t('form.email.name')"></el-input>
      </el-form-item>
      <el-form-item prop="name">
        <el-input v-model="form.name" :placeholder="$t('form.name.name')"></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input
          type="password"
          v-model="form.password"
          :placeholder="$t('form.password.name')"
          maxlength=20>
        </el-input>
      </el-form-item>
      <el-form-item prop="passwordConfirmation">
        <el-input
          type="password"
          v-model="form.passwordConfirmation"
          :placeholder="$t('form.passwordConfirmation.name')">
        </el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">
          {{ $t("views.signup.submitButton") }}
        </el-button>
      </el-form-item>
    </el-form>
    <p>
      {{ $t("views.signup.signInProposal") }}
      <el-button type="text" @click="showSignin()">
        {{ $t("components.header.signInButton") }}
      </el-button>.
    </p>
  </div>
</template>

<script>
import { PublicService } from '../services';

export default {
  name: 'signup',
  computed: {
    domainName() {
      return 'herd.anonbecon.com'
    }
  },
  methods: {
    showSignin() {
      this.$store.dispatch('showModal', 'signin');
    },
    onSubmit() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          this.$store.dispatch('signUp', this.form);
          this.$store.dispatch('closeModal', 'signup');
        }
      });
    },
  },
  data() {
    const validateConfirmation = (rule, value, callback) => {
      if (value !== this.form.password) {
        callback(new Error(this.$t('form.passwordConfirmation.incorrect')));
      } else {
        callback();
      }
    };

    const checkUsernameUniqueness = (rule, value, callback) => {
      PublicService.usernameAvailability(value).then(() => {
        callback();
      }).catch(() => {
        callback(new Error(this.$t('form.username.taken')));
      });
    };

    return {
      showUsernameLoadingIcon: false,
      form: {
        email: '',
        username: '',
        name: '',
        password: '',
        passwordConfirmation: '',
      },
      rules: {
        username: [
          { required: true, message: this.$t('form.username.required'), trigger: 'blur' },
          { min: 3, max: 20, message: this.$t('form.username.length', { min: 3, max: 20 }), trigger: 'blur' },
          { validator: checkUsernameUniqueness, trigger: 'blur' },
        ],
        email: [
          { required: true, message: this.$t('form.email.required'), trigger: 'blur' },
          { type: 'email', message: this.$t('form.email.incorrect'), trigger: 'blur' },
        ],
        name: [
          { required: true, message: this.$t('form.name.required'), trigger: 'blur' },
          { min: 3, max: 20, message: this.$t('form.name.length', { min: 3, max: 20 }), trigger: 'blur' },
        ],
        password: [
          { required: true, message: this.$t('form.password.required'), trigger: 'blur' },
          { min: 8, message: this.$t('form.password.length', { min: 8 }), trigger: 'blur' },
        ],
        passwordConfirmation: [
          { required: true, message: this.$t('form.passwordConfirmation.required'), trigger: 'blur' },
          { validator: validateConfirmation, trigger: 'blur' },
        ],
      },
    };
  },
};
</script>
