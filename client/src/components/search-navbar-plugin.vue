<template>
  <div class="navbar-plugin search-plugin">
    <el-input
      size="small"
      class="search-input"
      placeholder="Search herd"
      >
      <i
        class="el-icon-search el-input__icon"
        slot="suffix"
      />
    </el-input>
  </div>
</template>

<script>
export default {
  name: 'search-navbar-plugin',
  methods: {
  },
}
</script>

<style>
.search-input input{
  border: 0;
}
</style>
