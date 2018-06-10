<template>
  <el-header>
    <el-dialog :visible="modals.signin" @close="closeModal('signin')" customClass="dialog-custom">
      <login/>
    </el-dialog>
    <el-dialog :visible="modals.signup" @close="closeModal('signup')" customClass="dialog-custom">
      <signup/>
    </el-dialog>
    <el-row type="flex" class="row-bg header-wrapper" justify="center">
      <el-col :span="14">
        <div class="menu-right" v-if="session.user">
          <el-dropdown trigger="hover">
            <span class="el-dropdown-link userinfo-inner">{{ session.user.name }}</span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item @click.native="$router.push({ name: 'new-story'})">
                {{ $t('components.header.dropdown.newStory') }}
              </el-dropdown-item>
              <el-dropdown-item>Profile</el-dropdown-item>
              <el-dropdown-item>Settings</el-dropdown-item>
              <el-dropdown-item divided @click.native="logout">
                {{ $t('components.header.dropdown.signOut') }}
              </el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </div>
        <div class="menu-right" v-else>
          <el-button type="text" @click="showModal('signin')">
            {{ $t('components.header.signInButton') }}
          </el-button>
          <el-button type="primary" @click="showModal('signup')" plain>
            {{ $t('components.header.signUpButton') }}
          </el-button>
        </div>
      </el-col>
    </el-row>
  </el-header>
</template>

<script>
import { mapGetters } from 'vuex';
import Login from '../views/login';
import Signup from '../views/signup';

export default {
  name: 'app-header',
  computed: mapGetters({ session: 'session', modals: 'modals' }),

  components: {
    Login,
    Signup,
  },

  methods: {
    // TODO: Move this logic to shared customized modal
    showModal(modal) {
      this.$store.dispatch('showModal', modal);
    },
    closeModal(modal) {
      this.$store.dispatch('closeModal', modal);
    },
    logout() {
      this.$store.dispatch('signOut').then(() => {
        this.$router.replace({ path: '/' });
      });
    },
  },
};
</script>
<style lang="scss">
.header-wrapper {
  color: #333;
  line-height: 60px;
  .menu-right {
    text-align: right;
    padding-right: 35px;
    float: right;
    .userinfo-inner {
      cursor: pointer;
      // color:#fff;
      img {
        width: 40px;
        height: 40px;
        border-radius: 20px;
        margin: 10px 0px 10px 10px;
        float: right;
      }
    }
  }
}

</style>
