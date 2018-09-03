<template>
  <div>
    <h2>{{ $t('views.profile.settings.title') }}</h2>
    <div class="settings-block">
      <h3>Profile settings</h3>
      <ul class="ul-unstyled">
        <li>
          <h4>Avatar</h4>
          <div>
            <avatar
              :username="user.name"
              :inline="true"
              :src="session.avatar.url"
              :size="60"
            />
            <input
              ref="avatar"
              type="file"
              class="el-button el-button--text"
              v-on:change="handleUploadInput()"
            />
            <el-button
              @click="uploadAvatar"
              type="text"
              :disabled="disableSaveAvatarButton"
            >
              Save
            </el-button>
          </div>
        </li>
        <li>
          <h4>Name</h4>
          <el-input v-model="user.name" @change="updateUser"></el-input>
        </li>
        <li>
          <h4>Username</h4>
          <el-input v-model="user.username" disabled></el-input>
        </li>
        <li>
          <h4>Bio</h4>
          <el-input v-model="user.bio" @change="updateUser"></el-input>
        </li>
        <li>
          <h4>Email</h4>
          <el-input v-model="user.email" disabled>
            <i class="el-icon-edit el-input__icon" slot="suffix" @click="handleIconClick" />
          </el-input>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import Avatar from 'vue-avatar';
import { MediaService, UsersService } from '@/services';
import {
  CHANGE_SESSION,
} from '@/store/mutations.type';

export default {
  name: 'settings',
  components: { Avatar },
  computed: {
    ...mapGetters(['session']),
    disableSaveAvatarButton() {
      this.avatar == ''
    }
  },
  data() {
    return {
      avatar: '',
      user: this.$store.getters.session.user,
    }
  },
  methods: {
    handleIconClick(){
      console.log('click');
    },
    updateUser(){
      UsersService.put(session.user.id, this.user);
    },
    clearFileInput(){
      const input = this.$refs.avatar;
      input.type = 'text';
      input.type = 'file';
    },
    handleUploadInput(){
      this.avatar = this.$refs.avatar.files[0];
    },
    uploadAvatar(){
      let self = this;
      MediaService.createMedia(this.avatar.type)
        .then((res) => {
          MediaService.uploadFile(res.data.upload_url, this.avatar)
            .then(function(){
              self.clearFileInput();
              return UsersService.put(self.user.id, { media_id: res.data.data.id });
            }).then(res => {
              return MediaService.get(res.data.data.media_id)
            }).then(({data}) => {
              this.$store.commit(CHANGE_SESSION, { avatar: data.data });
            })
        .catch(function(){
              //TODO: Show notice
        });
      })
    }
  },
};
</script>

<style>

</style>
