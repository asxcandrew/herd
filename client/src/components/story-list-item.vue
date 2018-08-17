<template>
  <el-col class="underlined">
    <el-dialog
      title="Delete"
      :visible.sync="dialogVisible"
      width="30%"
      center>
      <span>Deleted stories are gone forever. Are you sure?</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="confirmDialog">Confirm</el-button>
      </span>
    </el-dialog>
    <h3>
      <router-link
        :to="{ name: 'edit-story', params: { id: story.id }}"
        class="unstyled-link"
      >{{ story.title }}</router-link>
    </h3>
    <span>
      <span class="substring last-edited">
        {{ $t('views.profile.stories.lastEdited', { date: lastUpdatedDate }) }}
      </span>
      <el-button type="text" @click="openDialog">
        {{ $t('views.profile.stories.deleteButton') }}
      </el-button>
      <el-button type="text" @click="$router.push({ name: 'edit-story', params: { id: story.id }})">
        {{ $t('views.profile.stories.editButton') }}
      </el-button>
    </span>
  </el-col>
</template>

<script>
import dateFormat from 'dateformat';
import { DELETE_STORY } from '@/store/actions.type';

export default {
  name: 'story-list-item',
  props: ['story'],
  computed: {
    lastUpdatedDate(){
      return dateFormat(this.story.updated_at, 'mmm d yyyy');
    },
  },
  data() {
    return {
      dialogVisible: false
    }
  },
  methods: {
    openDialog() {
      this.dialogVisible = true;
    },
    confirmDialog() {
      this.dialogVisible = false;
      this.deleteStory();
    },
    deleteStory() {
      this.$store.dispatch(DELETE_STORY, this.story.id);
    },
  },
};
</script>
<style>
.last-edited{
  margin-right: 10px;
}
</style>
