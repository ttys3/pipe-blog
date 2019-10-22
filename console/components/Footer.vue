<template>
  <footer class="footer">
    Powered by <a href="https://nanodm.net/" target="_blank">NanoBlog</a> {{$store.state.version}}
    Based on <a href="https://github.com/b3log/pipe" target="_blank">Pipe</a>  •
  </footer>
</template>

<style lang="sass">
  @import '~assets/scss/_variables'

  .footer
    padding: 17px 15px
    border-top: 1px solid $border-color
    font-size: 16px
    background-color: $white
</style>
