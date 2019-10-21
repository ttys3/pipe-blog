<template>
  <div class="console" id="particles">
    <div class="card login__content" ref="content">
      <form action="" class="login__form">

        <label for="username" class="login__label">username
        <input type="text" class="login__input" name="username" id="username" v-model="username" v-on:keyup.enter="login"/>
        </label>

        <label for="password" class="login__label">password
        <input type="password" class="login__input" name="password" id="password" v-model="password" v-on:keyup.enter="login"/>
        </label>

        <v-btn class="btn--small btn--info" @click="login">Login</v-btn>
      </form>
    </div>
  </div>
</template>

<script>
  import 'particles.js'
  import { initParticlesJS } from '~/plugins/utils'

  export default {
    data () {
      return {
        password: "",
        username: "",
      }
    },
    head () {
      return {
        title: this.$t('welcome', this.$store.state.locale) + ' - Pipe',
      }
    },
    methods: {
      toggleIntro () {
        this.$set(this, 'showIntro', !this.showIntro)
      },
      async doLogin () {
        const responseData = await this.axios.post(`/login`, {
          username: this.username,
          password: this.password,
        })
        if (responseData) {
          console.log(responseData)
        }
        return responseData
      },
      async login (event) {
        event.preventDefault()
        this.$store.commit('setSnackBar', {
          snackBar: true,
          snackMsg: this.$t('processing', this.$store.state.locale),
          snackModify: 'success',
        })
        try {
          const rs = await this.doLogin()
          console.log('rs: %o', rs)
          if (rs.code === 0) {
            this.$store.commit('setSnackBar', {
              snackBar: true,
              snackMsg: rs.msg,
              snackModify: 'success',
            })
            window.location.href = "/admin/"
          } else {
            this.$store.commit('setSnackBar', {
              snackBar: true,
              snackMsg: rs.msg,
            })
          }
        } catch (e) {
          console.log(e)
        }
      },
    },
    mounted () {
      initParticlesJS('particles')
    },
  }
</script>

<style lang="sass">
  .ft__12
    font-size: 12px !important
  .login
    &__intro
      text-align: left
      width: 300px
      margin: 0 auto
      ul
        margin-bottom: 0 !important
    &__space
      height: 10px
    &__checkbox
      margin: 0 20px
      color: #999
      a
        text-decoration: underline
        color: #67757c
    &__form
        padding: 2em
    &__input
        border-width: 1px
        height: 2em
        width: 15vw
        display: inline-block
    &__label
        display: block
        margin: 2em 5em

  @media (max-width: 470px)
    .login__checkbox
      line-height: 18px !important
</style>
