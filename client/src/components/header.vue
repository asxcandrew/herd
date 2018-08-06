<template>
  <el-header>
    <el-dialog v-if="modals.signin" :visible="modals.signin" @close="closeModal('signin')" customClass="dialog-custom">
      <login/>
    </el-dialog>
    <el-dialog v-if="modals.signup" :visible="modals.signup" @close="closeModal('signup')" customClass="dialog-custom">
      <signup/>
    </el-dialog>
    <el-row type="flex" class="row-bg header-wrapper" justify="center">
      <el-col :span="14">
        <el-col :span="3">
          <router-link to="/" class="logo unstyled-link">Herd</router-link>
        </el-col>
        <el-col :span="21">
          <el-row type="flex" justify="end">
            <el-col class="status-bar" :span="3">
              <i class="el-icon-loading" v-if="statusBar.isLoading"></i>
              <span>{{ statusBar.notice }}</span>
            </el-col>
            <el-col class="menu-right" v-if="session.user" :span="4">
              <el-dropdown trigger="hover">
                <span class="el-dropdown-link userinfo-inner">{{ session.user.name }}</span>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item @click.native="$router.push({ name: 'new-story'})">
                    {{ $t('components.header.dropdown.newStory') }}
                  </el-dropdown-item>
                  <el-dropdown-item @click.native="$router.push({ name: 'stories'})">
                    {{ $t('components.header.dropdown.stories') }}
                  </el-dropdown-item>
                  <el-dropdown-item>Bookmarks</el-dropdown-item>
                  <el-dropdown-item>Profile</el-dropdown-item>
                  <el-dropdown-item>Settings</el-dropdown-item>
                  <el-dropdown-item divided @click.native="logout">
                    {{ $t('components.header.dropdown.signOut') }}
                  </el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </el-col>
            <el-col class="menu-right" v-else :span="8">
              <el-button type="text" @click="showModal('signin')">
                {{ $t('components.header.signInButton') }}
              </el-button>
              <el-button type="primary" @click="showModal('signup')" plain>
                {{ $t('components.header.signUpButton') }}
              </el-button>
            </el-col>
          </el-row>
        </el-col>
      </el-col>
    </el-row>
  </el-header>
</template>

<script>
import { mapGetters } from 'vuex';
import Login from '@/views/login';
import Signup from '@/views/signup';
import { SIGN_OUT } from '@/store/actions.type';

export default {
  name: 'app-header',
  computed: mapGetters({ session: 'session', modals: 'modals', statusBar: 'statusBar' }),

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
      this.$store.dispatch(SIGN_OUT).then(() => {
        this.$router.replace({ path: '/' });
      });
    },
  },
};
</script>
<style lang="scss">
.header-wrapper {
  line-height: 60px;
  .logo {
    font-size: 20pt;
    font-weight: bold;
  }
  .status-bar {
    font-size: 9pt;
    color: #606266;
  }
  .menu-right {
    text-align: right;
    .userinfo-inner {
      cursor: pointer;
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
